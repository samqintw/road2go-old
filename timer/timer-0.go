package main

import (
	"time"
	"fmt"
)

func main()  {
	loopTimer := time.NewTimer(time.Second * 4)
	for {
		fmt.Println("Inside the infinite loop, waiting...")

		// blocking
		<-loopTimer.C
		break
	}
	fmt.Println("end infinite loop")
}
