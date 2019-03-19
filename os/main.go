package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func main() {
	// command-line arguments
	fmt.Println("command-line arguments:", os.Args)
	for i, arg := range os.Args {
		fmt.Println("command-line arguments ",i," : ", arg)
	}

	// file existence check
	_, err := os.Stat("i-dont-exist")
	fmt.Printf("%v\n", os.IsNotExist(err))

	// print current directory
	currDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current dir is %s\n", currDir)

	// Get Environment Variable
	fmt.Printf("Current $PATH is %s\n", os.Getenv("PATH"))

	// Get current username
	currUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Current User name is %s", currUser.Username)

	// Note: Platform dependant command
	cmd := exec.Command("ls", "-ltr")
	// Convenient wrapper
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)

	// Will not get called
	defer fmt.Println("Call Exit")
	os.Exit(0)
}