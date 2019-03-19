package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main()  {

	// log is thread safe, but fmt isn't
	log.Println("===== fmt vs. log ======")
	for i :=0; i<100;i++ {
		go log.Println("log ", i)
	}

	for i :=0; i<100;i++ {
		go fmt.Println("fmt ",i)
	}

	//
	log.Println("===== Log Flags ======")
	cf := log.Flags()
	cp := log.Prefix()
	log.Println("Current flag is ", cf)
	for i := 0; i <= 23; i++ {
		log.SetFlags(0)
		log.SetPrefix("")
		log.Print("Log Flag ", i)
		log.SetFlags(i)
		log.SetPrefix("    ")
		log.Println("")
	}

	log.SetFlags(cf)
	log.SetPrefix(cp)
	log.Println("\n", "===== Prefix ======")
	log.Println("Current prefix is ", cp)
	log.SetPrefix("myPrefix")
	log.Println("have my prefix")
	log.SetPrefix(cp)

	log.Println("\n===== Output======")
	log.Println("Discard")
	log.SetOutput(ioutil.Discard)
	log.Println("Discarding")

	log.SetOutput(io.Writer(os.Stdout))
	log.Println("Stdouting")

	log.SetOutput(io.Writer(os.Stderr))
	log.Println("Stderring")
	//log.SetOutput()

	log.Println("\n===== Fatal======")
	defer log.Println("Will not be logged")
	log.Fatal("exit!!!")

}
