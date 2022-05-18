package main

// Import necessary packages
import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"gopkg.in/yaml.v3"

	pb "example.com/go-tokenmgmt-grpc/tokenmgmt"
	"google.golang.org/grpc"
)

// Arrays to store token pointers and token IDs
var token_list []TokenData
var token_id_list []uint32
var errorlog *os.File
var logger *log.Logger
var port_nm string
var fs_timestamp time.Time

// TokenManagementServer definition
type TokenManagementServer struct {
	pb.UnimplementedTokenManagementServer
}

// Token Definition
type TokenData struct {
	mtx        sync.RWMutex
	id         uint32
	name       string
	low        uint64
	mid        uint64
	high       uint64
	partialval uint64
	finalval   uint64
	lst_tstmp  uint64
}

// Struct to read token data from YAML file
type TokenAccess struct {
	Token   int    `yaml:"token"`
	Writer  string `yaml:"writer"`
	Readers string `yaml:"readers"`
}

// Hash concatentates a message and a nonce and generates a hash value.
func Hash(name string, nonce uint64) uint64 {
	hasher := sha256.New()
	hasher.Write([]byte(fmt.Sprintf("%s %d", name, nonce)))
	return binary.BigEndian.Uint64(hasher.Sum(nil))
}

// Function to retrieve a specific token's data from the YAML file
func yaml_data_retriever(token_id int) TokenAccess {

	// Open the YAML file and read the file as byte
	yfile, err := ioutil.ReadFile("yaml_final.yml")
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Decoder object to parse YAML structure
	decoder := yaml.NewDecoder(bytes.NewBufferString(string(yfile)))
	for {
		var sample TokenAccess
		err = decoder.Decode(&sample)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		// If the current token information is the one we need, return the information
		if sample.Token == token_id {
			return sample
		}
	}
	logger.Println("Decoded everything.")
	return TokenAccess{}
}

// Check if a token with a given ID already exists
func is_exists(id_val uint32) (int, bool) {

	tmp := -1

	// Iterate through the token ID list and return the index if present
	for ind, val := range token_id_list {
		if val == id_val {
			return ind, true
		}
	}
	// If not present, return -1 and False as result
	return tmp, false
}

// Print all token IDs except for one specific token
func print_all_tokens(ind_val int) {

	// Create an empty string to store the IDs of the tokens
	logger.Println("Other token IDs:")
	other_tokens := ""

	for ind := 0; ind < len(token_id_list); ind++ {
		if ind != int(ind_val) {
			other_tokens += strconv.Itoa(int(token_id_list[ind])) + " "

		}
	}

	// Print the information in the log file
	logger.Println(other_tokens)
}

// Print current token information
func print_current_token(ind_val int) {
	logger.Println("Current token data:")
	logger.Printf("ID: %d, Name: %s, Low: %d, Mid: %d, High: %d\n", token_list[ind_val].id, token_list[ind_val].name, token_list[ind_val].low, token_list[ind_val].mid, token_list[ind_val].high)
}

// Function to assist fail-silent emilation
func fail_silent_check(token_id uint32) bool {

	// Function works only for the token 1020 on server 65000
	_, status := is_exists(token_id)
	if status && (port_nm == "65000") && (token_id == 1020) {
		curr_time := time.Now()
		diff := curr_time.Sub(fs_timestamp).Seconds()
		if diff > 10 {
			logger.Println("10 second deadline reached, fail-silent behaviour enforced.")
		}
		return diff > 10
	}
	return false
}

// Function to get the list of reader ports
func get_port_list(token_id uint32) []string {

	var readers []string
	var port_list []string

	// Fetch the current token information from the file
	token_info := yaml_data_retriever(int(token_id))

	// Update the list of readers
	if token_info.Token != 0 {
		readers = strings.Split(token_info.Readers, " ")
	}

	// Iterate through the readers to update the list of reader ports
	for _, element := range readers {
		port := strings.Index(element, ":")
		if element[len(element)-1:] == "," {
			port_list = append(port_list, element[port+1:len(element)-1])
		} else {
			port_list = append(port_list, element[port+1:])
		}
	}

	// Return the list of ports
	return port_list
}

// Support function that performs Read-Impose Write-Majority
func get_finalvals(ch chan []uint64, element string, token_id uint32, lst_tstmp uint64) {

	// Connect to the node whose port is given in "element"
	addr := "localhost:" + element
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	// 3 second context timeout
	c := pb.NewTokenManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Make the RIWMTest call to enforce Read-Impose Write-Majority
	res, err := c.RIWMTest(ctx, &pb.RIWMInput{Id: token_id, LstTstmp: lst_tstmp})

	// Store the data and write back to channel if no errors
	if err == nil {
		var data []uint64
		data = append(data, res.LstTstmp)
		data = append(data, res.Finalval)
		logger.Println("Response from the server running on port", element, "in the quorum:", data)
		ch <- data
	}
}

// Create RPC call definition
func (s *TokenManagementServer) Create(ctx context.Context, in *pb.CreateInput) (*pb.SuccessStatus, error) {
	var msg string

	// Check if the token already exists
	if fail_silent_check(in.GetId()) {
		time.Sleep(10 * time.Second)
		return &pb.SuccessStatus{Msg: msg}, nil
	} else {
		logger.Println("-----------------------------------------------------------")
		logger.Println("Operation: Create")
		msg = "Error encountered."

		ind, res := is_exists(in.GetId())

		// If the token does not exist, create the token
		if !(res) {
			tmp := 0
			token1 := &TokenData{
				id:         in.GetId(),
				name:       "",
				low:        uint64(tmp),
				mid:        uint64(tmp),
				high:       uint64(tmp),
				partialval: uint64(tmp),
				finalval:   uint64(tmp),
				lst_tstmp:  in.GetLstTstmp(),
			}

			// Lock the newly created token until it is added to the list
			token1.mtx.Lock()
			token_list = append(token_list, *token1)
			token_id_list = append(token_id_list, token1.id)
			token_list[len(token_list)-1].mtx.Unlock()
			print_current_token(len(token_list) - 1)
			msg = "Token created successfully."
			ind = len(token_list) - 1

			// Update the timestamp for the token named 1020
			if token1.id == 1020 {
				fs_timestamp = time.Now()
				logger.Println("Timestamp is set.")
			}
		} else {
			msg = "Token already exists."
			logger.Println(msg)
		}

		// If the source of the Write is client, then the writer has to update the readers
		if in.Source == "client" {
			port_list := get_port_list(in.GetId())
			for _, element := range port_list {

				// If the array element is not the same as the current server
				if port_nm != element {

					// Create the reader servers
					logger.Println("Sending Create RPC call to server running on port", element)
					cmd := exec.Command("servercode/servercode", "-port", element)
					err := cmd.Start()
					if err != nil {
						panic(err)
					}
					addr := "localhost:" + element
					conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						log.Fatalf("Could not connect: %v", err)
					}
					defer conn.Close()

					// Get context and set a 5 second timeout
					c := pb.NewTokenManagementClient(conn)
					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					// Send the Create RPC call to the reader to update its token
					res, _ := c.Create(ctx, &pb.CreateInput{Id: in.GetId(), Source: "writer", LstTstmp: in.GetLstTstmp()})
					msg = res.Msg
					logger.Println("Response from the server:", res.Msg)
				}
			}
		}
		print_all_tokens(ind)
		return &pb.SuccessStatus{Msg: msg}, nil
	}
}

// Read RPC call definition
func (s *TokenManagementServer) Read(ctx context.Context, in *pb.ReadInput) (*pb.ResultRead, error) {
	var tmp uint64
	tmp = 0

	// Check if the current token satisfies the fail-silent conditions
	if fail_silent_check(in.GetId()) {
		time.Sleep(10 * time.Second)
		return &pb.ResultRead{Finalval: tmp}, nil
	} else {
		logger.Println("-----------------------------------------------------------")
		logger.Println("Operation: Read")
		logger.Println(in.GetId())

		// Check if the token already exists
		ind, res := is_exists(in.GetId())
		// If the token exists
		if res {

			// Check if the write operation is already performed
			if token_list[ind].partialval != 0 {

				// Get the reader port lists and track the current token's finalval
				port_list := get_port_list(in.GetId())
				tmp = token_list[ind].finalval

				// Initialize a channel for the Read-Impose Write-Majority results and spawn the goroutines
				ch := make(chan []uint64)
				for element := range port_list {
					if port_list[element] != port_nm {
						go get_finalvals(ch, port_list[element], token_list[ind].id, token_list[ind].lst_tstmp)
					}
				}

				// Read from the channel at least until a majority of responses have been received
				ctr := 1
				for res := range ch {
					ctr = ctr + 1

					// Update the current token information if it is outdated
					if res[0] > token_list[ind].lst_tstmp {
						token_list[ind].lst_tstmp = res[0]
						token_list[ind].finalval = res[1]
						tmp = res[1]
					}
					if len(port_list) == 2 {
						if ctr == 2 {
							print_current_token(ind)
							print_all_tokens(ind)
							return &pb.ResultRead{Finalval: tmp}, nil
						}
					} else {
						if ctr > len(port_list)/2 {
							print_current_token(ind)
							print_all_tokens(ind)
							return &pb.ResultRead{Finalval: tmp}, nil
						}
					}
				}
				print_current_token(ind)
			} else {
				// If the write operation is not yet performed
				logger.Println("Token values are not set.")
			}
		} else {
			// If the token is yet to be created
			logger.Println("Token does not exist.")
		}
		print_all_tokens(ind)
		return &pb.ResultRead{Finalval: tmp}, nil
	}
}

// Write RPC call definition
func (s *TokenManagementServer) Write(ctx context.Context, in *pb.WriteInput) (*pb.ResultWrite, error) {
	var tmp uint64
	tmp = 0
	if fail_silent_check(in.GetId()) {
		time.Sleep(10 * time.Second)
		return &pb.ResultWrite{Partialval: tmp}, nil
	} else {
		logger.Println("-----------------------------------------------------------")
		logger.Println("Operation: Write")
		// Check if the token already exists
		ind, res := is_exists(in.GetId())
		// If the token exists
		if res {

			// Lock the token because data is about to be written in the token
			if token_list[ind].lst_tstmp <= in.GetLstTstmp() {

				token_list[ind].mtx.Lock()
				token_list[ind].name = in.GetName()
				token_list[ind].low = in.GetLow()
				token_list[ind].mid = in.GetMid()
				token_list[ind].high = in.GetHigh()
				token_list[ind].lst_tstmp = in.GetLstTstmp()
				val := Hash(token_list[ind].name, token_list[ind].low)
				for i := token_list[ind].low + 1; i < token_list[ind].mid; i++ {
					tmp = Hash(token_list[ind].name, i)
					if tmp < val {
						val = tmp
					}
				}
				// tm := uint64(0)
				token_list[ind].partialval = val
				val = Hash(token_list[ind].name, token_list[ind].mid)
				for i := token_list[ind].mid + 1; i < token_list[ind].high; i++ {
					tmp = Hash(token_list[ind].name, i)
					if tmp < val {
						val = tmp
					}
				}
				token_list[ind].finalval = val
				tmp = val
				// Release the lock now that the data is updated
				token_list[ind].mtx.Unlock()
				print_current_token(ind)
			}
		} else {
			logger.Println("Token does not exist.")
		}

		port_list := get_port_list(in.GetId())

		// logger.Println(in.Source, port_list, in.Source == "client")
		if in.Source == "client" {
			for _, element := range port_list {
				// logger.Println(in.Source, element, port_nm, len(element), len(port_nm), element == port_nm)
				if port_nm != element {
					addr := "localhost:" + element
					// logger.Println("this is a writer", addr)
					conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
					if err != nil {
						log.Fatalf("Could not connect: %v", err)
					}
					defer conn.Close()

					// Get context and set a 10 second timeout
					c := pb.NewTokenManagementClient(conn)
					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()
					logger.Println("Sending Write RPC call to server running on port", element)
					res, _ := c.Write(ctx, &pb.WriteInput{Id: in.GetId(), Name: in.GetName(), Low: in.GetLow(), Mid: in.GetMid(), High: in.GetHigh(), Source: "writer", LstTstmp: in.GetLstTstmp()})
					tmp = res.Partialval
					logger.Println("Response from the server:", res.Partialval)
				}
			}
		}
		print_all_tokens(ind)
		return &pb.ResultWrite{Partialval: tmp}, nil
	}
}

// Drop RPC call definition
func (s *TokenManagementServer) Drop(ctx context.Context, in *pb.CreateInput) (*pb.SuccessStatus, error) {
	var msg string
	if fail_silent_check(in.GetId()) {
		time.Sleep(10 * time.Second)
		return &pb.SuccessStatus{Msg: msg}, errors.New("token fail silent on server")
	} else {
		logger.Println("-----------------------------------------------------------")
		logger.Println("Operation: Drop")
		// Check if the token already exists
		ind, res := is_exists(in.GetId())
		// If the token exists
		var del_tstmp uint64
		if res {
			// Drop the token by removing it from the token_list container
			print_current_token(ind)
			tmp := token_list[ind]
			del_tstmp = tmp.lst_tstmp
			tmp.mtx.Lock()
			token_list = append(token_list[:ind], token_list[ind+1:]...)
			tmp.mtx.Unlock()
			token_id_list = append(token_id_list[:ind], token_id_list[ind+1:]...)
			logger.Println("Token deleted successfully.")
			msg = "Token deletion successful."
		} else {
			// Token does not exist
			logger.Println("Token does not exist.")
			msg = "Token does not exist, deletion impossible."
		}
		if in.Source == "client" {
			port_list := get_port_list(in.GetId())
			for element := range port_list {
				if port_list[element] != port_nm {
					logger.Println("Sending Drop RPC call to server running on port", port_list[element])
					go func(element string, token_id uint32, lst_tstmp uint64) {
						addr := "localhost:" + element
						conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
						if err != nil {
							log.Fatalf("Could not connect: %v", err)
						}
						defer conn.Close()
						c := pb.NewTokenManagementClient(conn)
						ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
						defer cancel()
						res, err := c.Drop(ctx, &pb.CreateInput{Id: token_id, LstTstmp: lst_tstmp, Source: "writer"})
						if err == nil {
							data := res.Msg
							logger.Println("Response from the server:", data)
						}
					}(port_list[element], in.GetId(), del_tstmp)

				}
			}

		}
		print_all_tokens(-1)
		return &pb.SuccessStatus{Msg: msg}, nil
	}
}

// RIWMTest RPC call definition
func (s *TokenManagementServer) RIWMTest(ctx context.Context, in *pb.RIWMInput) (*pb.RIWMOutput, error) {
	if fail_silent_check(in.GetId()) {
		time.Sleep(10 * time.Second)
		return &pb.RIWMOutput{LstTstmp: 0, Finalval: 0}, errors.New("token fail silent on server")
	} else {
		ind, res := is_exists(in.GetId())
		if res {
			return &pb.RIWMOutput{LstTstmp: token_list[ind].lst_tstmp, Finalval: token_list[ind].finalval}, nil
		} else {
			return &pb.RIWMOutput{LstTstmp: 0, Finalval: 0}, errors.New("token does not exist")
		}
	}
}

// Main function
func main() {
	// Flag to parse the port argument

	portflag := flag.Int("port", 50051, "port to connect to")
	flag.Parse()
	port_nm = strconv.Itoa(*portflag)
	fl_nm := "logs/log_" + strconv.Itoa(*portflag) + ".log"
	errorlog, _ = os.OpenFile(fl_nm, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer errorlog.Close()
	logger = log.New(errorlog, "applog: ", log.Lshortfile|log.LstdFlags)
	// Connect to the port mentioned in the command line arguments
	port := ":" + strconv.Itoa(*portflag)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Register new server and have it running on the port
	s := grpc.NewServer()
	pb.RegisterTokenManagementServer(s, &TokenManagementServer{})
	log.Printf("\nServer started. Listening at the following address: %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
