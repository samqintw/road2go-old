package main

import "fmt"

func main()  {
	defer fmt.Println("World 1")
	defer fmt.Println("World 2")
	defer func() {
		fmt.Println("recovery func 1")
		if err := recover(); err!=nil {
			fmt.Println(err)
			fmt.Println("go recovery")
		}
	}()
	//defer fmt.Println("World 3")
	//defer func() {
	//	fmt.Println("recovery func 2")
	//	if err := recover(); err!=nil {
	//		fmt.Println(err)
	//		fmt.Println("go recovery")
	//	}
	//}()
	//defer fmt.Println("World 4")
	fmt.Println("Hello")

	gopanic("go panic1")
	fmt.Println("Hello 2")
	//gopanic("go panic2")
}

func gopanic(msg string)  {
	panic(msg)
}
