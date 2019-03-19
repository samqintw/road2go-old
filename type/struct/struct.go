package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age int
	Address string
}

func main()  {
	// return
	d1 := person{
		Name:"somin chin",
		Age:38,
		Address:"Taiwan",
	}
	fmt.Println(d1)
	fmt.Println(reflect.TypeOf(d1))

	d2 := &person{
		Name:"somin chin",
		Age:38,
		Address:"Taiwan",
	}
	fmt.Println(d2)
	fmt.Println(reflect.TypeOf(d2))

	d3 := new(person) // this equals to &person{}
	d3.Name = "somin chin"
	fmt.Println(d3)
	fmt.Println(reflect.TypeOf(d3))

}