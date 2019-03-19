package main

import (
	"encoding/json"
	"fmt"
)

type Obj struct {
	Name string
	Message string `json:"message"` // rename
	Author string `json:"-"` // do not output the field
	Date string `json:",omitempty"` // do not output the field if the value is empty
	age string
}

func main()  {
	data, err := json.Marshal(Obj{Name: "smchin1", Message: "helloworld", age: "20"})
	if err != nil {
		panic("Oops")
	}
	fmt.Println(string(data))

	data, err = json.Marshal(Obj{Name: "smchin2", Message: "helloworld", Author:"somin", Date:"2018"})
	if err != nil {
		panic("Oops")
	}
	fmt.Println(string(data))
}
