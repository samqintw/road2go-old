package main

import (
	"flag"
	"fmt"
)

var (
	name = flag.String("name", "smchin", "--name=smchin")
)

//go:generate go run main.go --name=somin
func main()  {
	fmt.Println("default name: ", *name)
	flag.Parse()
	fmt.Println("parsed name: ", *name)
	flag.Lookup("name").Value.Set("sung-ming")
	fmt.Println("assigned name: ", *name)
}
