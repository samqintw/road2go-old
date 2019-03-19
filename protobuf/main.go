package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"myGo/protobuf/example"
	"net"
	"strings"
)

func main() {
	op := flag.String("op", "s", "s for server, c for client!")
	flag.Parse()
	fmt.Println("flag: ", *op)
	switch strings.ToLower(*op) {
	case "c":
		runClient()
	case "s":
		runServer()
	}
}

func runServer()  {
	listener, err := net.Listen("tcp", ":8282")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		fmt.Println("Listening...")
		conn, err := listener.Accept()
		fmt.Println("Connection Accepted")
		if err != nil {
			log.Fatal(err)
		}
		go func(conn net.Conn) {
			defer conn.Close()
			fmt.Println("processing Conn")
			data, err := ioutil.ReadAll(conn)
			if err != nil {
				return
			}
			fmt.Println("Received data: ", data)
			a := &example.Person{}
			proto.Unmarshal(data, a)
			fmt.Println("Received data: ", a)
		}(conn)
	}
}

func runClient()  {
	a := &example.Person{
		Id: 12,
		Name: "smchin",

	}
	data, err := proto.Marshal(a)
	if err != nil {
		log.Fatal(err)
		return
	}
	SendData([]byte(data))
}

func SendData(data []byte)  {
	fmt.Println("Sending Data ...")
	conn, err := net.Dial("tcp", "127.0.0.1:8282")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	conn.Write(data)
}