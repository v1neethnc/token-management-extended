package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type TokenAccess struct {
	Token   int    `yaml:"token"`
	Writer  string `yaml:"writer"`
	Readers string `yaml:"readers"`
}

func main() {

	yfile, err := ioutil.ReadFile("test_yaml.yml")
	if err != nil {
		log.Fatal(err)
	}
	data := TokenAccess{}
	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(data.Token, "bruh", data.Writer, "bruhruh", data.Readers)
	fmt.Printf("%T, %T, %T\n", data.Token, data.Writer, data.Readers)
	temp := strings.Split(data.Readers, ",")
	var readers []string
	for i := 0; i < len(temp); i++ {
		readers = append(readers, strings.Trim(temp[i], " s"))
	}
	fmt.Println(readers[0:1], readers[1:2], readers[2:])
}
