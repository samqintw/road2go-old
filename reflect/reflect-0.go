package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num = 1.2345

	fmt.Println("type: ", reflect.TypeOf(num))
	fmt.Println("value: ", reflect.ValueOf(num))
}