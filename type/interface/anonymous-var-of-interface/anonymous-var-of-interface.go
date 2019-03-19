package anonymous_var_of_interface

import "fmt"

// type alias
type Helper = interface {
	Help1(in string) string
	Help2(in string) string
}

type HelpString string

func (hs HelpString) Help1(in string) string  {
	return string("[help1]" + string(hs) + ", " + in)
}

func (hs HelpString) Help2(in string) string  {
	return string("[help2]" + string(hs) + ", " + in)
}

func main()  {
	var helpString = HelpString("help me")

	var explicit1 = interface{
		Help1(in string) string // single line is not allowed
		Help2(in string) string
	}.Help1(helpString, "smchin1")
	fmt.Println(explicit1)

	var explicit2 = interface{Help1(in string) string; Help2(in string) string}.Help1(helpString, "smchin2")
	fmt.Println(explicit2)

	var explicit3 = interface{Help1(in string) string}.Help1(helpString, "smchin3")
	fmt.Println(explicit3)

	// declare an anonymous interface of Helper and assign h to it,
	// then calling help2 method with param "somin1"
	var my1 = Helper.Help2(helpString /* assign h to anonymous interface var */, "somin1" /* param for help2 method */)
	fmt.Println(my1)

	// declare iHelper as Helper interface
	var iHelper Helper = helpString
	var my2 = iHelper.Help2("somin2")
	fmt.Println(my2)
}