package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.Gosched()
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(1)
	go go1(&wg)
	wg.Add(1)
	go go2(&wg)

	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
}

func go1(wg *sync.WaitGroup) {
	fmt.Println(">>> go1")
	for i := 0; i < 20; i++ {
		fmt.Println("go1: ", i)
	}
	fmt.Println("<<< go2")
	wg.Done()
}

func go2(wg *sync.WaitGroup) {
	fmt.Println(">>> go2")
	for i := 0; i < 10; i++ {
		fmt.Println("go2: ", i)
	}
	fmt.Println("<<< go2")
	wg.Done()
}
