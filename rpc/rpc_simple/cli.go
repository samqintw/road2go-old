package main

import (
	"fmt"

	"github.com/samqintw/road2go/rpc/rpc_simple/client"
	"github.com/samqintw/road2go/rpc/rpc_simple/server"
)

func main() {
	go server.StartServer()

	c := client.CreateClient()
	defer c.Close()

	reply := c.PerformRequest()
	fmt.Println(reply.Message)
}