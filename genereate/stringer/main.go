package main

import (
	"fmt"
	"myGo/genereate/stringer/painkiller"
)

//go:generate stringer -type=Pill painkiller

func main()  {
	fmt.Println(painkiller.Aspirin)
}