package main

import (
	"fmt"

	"github.com/samqintw/road2go/rpc/rpc_http/client"
	"github.com/samqintw/road2go/rpc/rpc_http/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := client.PerformRequest(c)

	fmt.Println(reply.Message)
}