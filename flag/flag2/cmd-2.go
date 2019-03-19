package flag2

import (
	"flag"
	"fmt"
	"regexp"
	"log"
)

const UsernameRegex string =`^@?(\w){1,15}$`

func main()  {
	var userNameInput string
	flag.StringVar(&userNameInput, "username", "Gopher", "The name of Gopher")
	flag.Parse()
	fmt.Println("Validation Checker!")
	fmt.Println("Checking Syntax for username, \"", userNameInput, "\" result in: ", CheckUserNameSyntax(userNameInput))
}

func CheckUserNameSyntax(userName string) bool {
	validationResult := false
	r, err := regexp.Compile(UsernameRegex)
	if err != nil {
		log.Fatal(err)
	}
	validationResult = r.MatchString(userName)
	return validationResult
}