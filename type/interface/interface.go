package main

import (
	"fmt"
)


type myiface interface {
	SayHello()
	Say(msg string)
	Increase()
	GetI() int
}

type testConcreteImpl struct {
	i int
}

// use global anonymous func to embed method to struct
// as known, receiver syntax
func (tst testConcreteImpl) SayHello() {
	fmt.Printf("&tst=%p\n", &tst)
	//debug.PrintStack()
}

func (tst testConcreteImpl) Say(msg string) {
	fmt.Println(msg)
	//debug.PrintStack()
}

func (tst *testConcreteImpl) Increase() {
	fmt.Printf("*tst=%p\n", tst)
	fmt.Printf("&tst=%p\n", &tst)
	(*tst).i++
}

func (tst *testConcreteImpl) GetI() int {
	return (*tst).i
}

type myEmbedding struct {
	*testConcreteImpl
}

func myfunc(str string) string {
	fmt.Println("myfunc", str)
	return "smchin"
}


func main() {
	tc := testConcreteImpl{}
	fmt.Printf("&tc = %p\n", &tc)
	var iface myiface = &tc
	iface.SayHello()
	fmt.Printf("i is %d\n", tc.i)
	iface.Increase()
	fmt.Printf("i is %d\n", tc.i)
	iface.Increase()
	fmt.Printf("i is %d\n", iface.GetI())

	var mi interface{} = myfunc
	fmt.Printf("%T\n", mi)
	if v, ok := mi.(func(string) string); ok {
		fmt.Println("func assertion")
		v("somin")
	}

	var myiface1 interface{} = iface
	fmt.Printf("%T\n", myiface1)
	if v, ok := myiface1.(myiface); ok {

	}

}
