package main

import (
	"flag"
	"fmt"
)

//go:generate go run cmd-1.go --name=somin
func main()  {
	var gopherName string
	flag.StringVar(&gopherName, "name", "Gopher", "The name of Gopher")
	fmt.Println("[before Parse] Hello ", gopherName, "!")
	flag.Parse()
	fmt.Println("[after Parse] Hello ", gopherName, "!")
	flag.Parse()
	fmt.Println("[after Parse] Hello ", gopherName, "!")
	fmt.Println("Hello ", gopherName, "!")
}