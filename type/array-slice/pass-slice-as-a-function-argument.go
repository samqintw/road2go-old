package main

import (
	"fmt"
	)


func init()  {

}

func modifyValue(s []int)  {
	s[1] = 3
	fmt.Printf("In modifyValue: s is %v, address is %p\n", s, s)
}

func addValue(s []int) {
	fmt.Printf("In addValue: s is %v, address is %p\n", s, s)
	ss := append(s, 3)
	fmt.Printf("In addValue: s is %v, address is %p\n", ss, ss)
}

func addValueByPoint(s *[]int) {
	*s = append(*s, 3)
	(*s)[1] = 5
	fmt.Printf("In addValue: s is %v\n", s)
}

func main() {
	s := []int{1, 2}
	fmt.Printf("In main, s address is %p\n", s)
	// s is slice
	fmt.Printf("In main, s is %v\n", s)
	modifyValue(s)
	fmt.Printf("In main, s is %v\n", s)
	addValue(s)
	fmt.Printf("In main, s is %v\n", s)
	addValueByPoint(&s)
	fmt.Printf("In main, s is %v\n", s)
}