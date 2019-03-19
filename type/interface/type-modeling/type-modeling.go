package main

import (
	"fmt"
)

type WebController struct {}

func (wc *WebController)GetName () string {
	return "Web Controller"
}

type Indexer interface {
	Index()
}

// Anonymous type embedding
type AppController struct {
	*WebController
	Indexer
}

type IndexString string

func (hs IndexString) Index()  {
	fmt.Println("Index Page " + hs)
}

func main() {
	ac := new(AppController)
	fmt.Println(ac.WebController.GetName())
	// shorthand
	fmt.Println(ac.GetName())
	//ac.Index() // go panics, because Indexer is nil
	ac.Indexer = IndexString("somin")
	ac.Index()

	ac = &AppController{new(WebController), IndexString("scmhin")}
	ac.Index()
}