package client

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/samqintw/road2go/rpc/rpc_simple/contract"
)

const port = 1234

type MyClient struct {
	*rpc.Client
}

func CreateClient() *MyClient {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return &MyClient{client}
}

func (my MyClient) PerformRequest() contract.HelloWorldResponse {
	args := &contract.HelloWorldRequest{Name: "World Smchin"}
	var reply contract.HelloWorldResponse

	err := my.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}

	return reply
}