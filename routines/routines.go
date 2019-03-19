package main

import (
	"fmt"
	"time"
)

func main()  {
	ic := make(chan int)
	go periodicSend(ic)
	for received := range ic {
		fmt.Println(received)
	}
	fmt.Println("closed")
}

func periodicSend(ic chan int) {
	amount := 0;
	for amount < 10 {
		amount = amount + 1
		ic <- amount
		time.Sleep(1*time.Second)
	}
	close(ic)
}
