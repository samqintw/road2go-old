package main

import "fmt"

// type alias
type iHelper = interface {
	Help() string
}

type HelpString string

func (hs HelpString) Help() string  {
	return string(hs)
}

type UnHelpString struct {}

func (uhs *UnHelpString) Help() string  {
	return "I cannot help you!"
}

// Compile time check
var _ = iHelper(HelpString(""))

func main()  {
	var helpers = []iHelper{
		HelpString("Help me again!"), // this is a HelpString type
		&UnHelpString{}, // this is a struct
	}

	fmt.Println(helpers)

	for _, helper := range helpers {
		fmt.Println(helper.Help())
	}

	var h2 interface{} = HelpString("Please help me more!")
	n, ok := h2.(iHelper) // # Type Assertions
	fmt.Println(n, ok)

	// it convert interface to a string type,
	// only can convert an  interface to another an interface
	var _ = h2.(string) // go panic at runtime
}