package main

import (
	"fmt"
	"runtime"
)

func showNumber (i int) {
	fmt.Println(i)
}

func main() {

	for i := 0; i < 10; i++ {
		go showNumber(i)
	}
	// to yield processor
	runtime.Gosched()
	fmt.Println("Haha")
}