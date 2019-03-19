package main

import "fmt"

func main() {
	// init
	x := make(map[string]int)

	/*
	var x map[string]int = make(map[string]int)
	x1 := map[string]int {
		"third": 3,
		"forth": 4,
	}
	 */

	// access
	x["first"] = 1
	x["second"] = 2

	if v, ok := x["first"]; ok {
		fmt.Println(v)
	}

	// delete
	delete(x, "first")
	fmt.Println(x)
}
