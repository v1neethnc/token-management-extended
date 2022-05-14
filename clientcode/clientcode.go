package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"

	pb "example.com/go-tokenmgmt-grpc/tokenmgmt"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

type TokenAccess struct {
	Token   int    `yaml:"token"`
	Writer  string `yaml:"writer"`
	Readers string `yaml:"readers"`
}

func yaml_data_retriever(token_id int) TokenAccess {
	yfile, err := ioutil.ReadFile("yaml_final.yml")
	if err != nil {
		log.Fatal(err)
	}
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
		if sample.Token == token_id {
			return sample
		}
	}
	fmt.Println("Decoded everything.")
	return TokenAccess{}
}

func main() {

	// Define the different flags required
	mastercmd := flag.NewFlagSet("test", flag.ExitOnError)
	idptr := mastercmd.Int("id", 0, "id of token")
	nameptr := mastercmd.String("name", "foobar", "name of token")
	// hostptr := mastercmd.String("host", "localhost", "host to connect to")
	// portptr := mastercmd.Int("port", 50051, "port to connect to")
	lowptr := mastercmd.Int("low", 0, "low value of token")
	midptr := mastercmd.Int("mid", 0, "mid value of token")
	highptr := mastercmd.Int("high", 0, "high value of token")
	// dropptr := mastercmd.Int("drop", 0, "id of token")

	var writer string
	// var id_val int
	var readers []string

	// Parse the command line arguments
	if os.Args[1] == "-drop" {
		mastercmd.Parse(os.Args[1:])
	} else {
		mastercmd.Parse(os.Args[2:])
	}
	token_info := yaml_data_retriever(*idptr)
	if token_info.Token != 0 {
		// id_val = token_info.Token
		writer = token_info.Writer
		readers = strings.Split(token_info.Readers, " ")
	}

	// Switch case to call different RPCs depending on the command line arguments
	switch os.Args[1] {
	case "-create":
		writer_port := strings.Index(writer, ":")
		// fmt.Println(writer[writer_port+1:])
		cmd := exec.Command("go", "run", "servercode/servercode.go", "-port", writer[writer_port+1:])
		// filename := "logs/" + writer[writer_port+1:] + "_log.txt"
		// outfile, err := os.Create(filename)
		// if err != nil {
		// 	panic(err)
		// }
		// defer outfile.Close()
		// cmd.Stdout = outfile

		err := cmd.Start()
		if err != nil {
			panic(err)
		}
		// wrt := "localhost:" + writer[writer_port+1:]
		// fmt.Println(wrt)
		conn, err := grpc.Dial(writer, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()

		// Get context and set a 10 second timeout
		c := pb.NewTokenManagementClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		// fmt.Println(c, ctx)
		defer cancel()
		res, _ := c.Create(ctx, &pb.CreateInput{Id: uint32(*idptr), Source: "client"})
		log.Println("\nResponse from the server:", res.Msg)

	case "-read":
		rand.Seed(time.Now().UnixNano())
		var reader_ports []string
		for _, element := range readers {
			port := strings.Index(element, ":")
			if element[len(element)-1:] == "," {
				reader_ports = append(reader_ports, element[port+1:len(element)-1])
			} else {
				reader_ports = append(reader_ports, element[port+1:])
			}
		}
		random_ind := rand.Intn(len(reader_ports))
		reader := reader_ports[random_ind]
		fmt.Println(reader, random_ind, reader_ports)
		rdr := "127.0.0.1:" + reader
		conn, err := grpc.Dial(rdr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()
		c := pb.NewTokenManagementClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, err := c.Read(ctx, &pb.ReadInput{Id: uint32(*idptr)})
		if err != nil {
			panic(err)
		} else {
			if res.Finalval == 0 {
				log.Println("\nResponse from the server: Token does not exist.")
			} else {
				log.Println("\nResponse from the server:", res.Finalval)
			}
		}

	case "-write":
		ind := strings.Index(writer, ":")
		port := writer[ind+1:]
		fmt.Println(port)
		conn, err := grpc.Dial(writer, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		defer conn.Close()

		// Get context and set a 10 second timeout
		c := pb.NewTokenManagementClient(conn)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		res, _ := c.Write(ctx, &pb.WriteInput{Id: uint32(*idptr), Name: *nameptr, Low: uint64(*lowptr), Mid: uint64(*midptr), High: uint64(*highptr), Source: "client"})
		if res.Partialval == 0 {
			log.Println("\nResponse from the server: Token does not exist.")
		} else {
			log.Println("\nResponse from the server:", res.Partialval)
		}

		// case "-drop":
		// 	res, _ := c.Drop(ctx, &pb.TokenData{Id: uint32(*dropptr), Host: *hostptr, Port: uint32(*portptr)})
		// 	log.Println("\nResponse from the server:", res.Msg)
	}
}
