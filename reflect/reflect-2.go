package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) ReflectCallFunc() {
	fmt.Println("Allen.Wu ReflectCallFunc")
}

func main() {

	user := User{1, "Allen.Wu", 25}

	DoFiledAndMethod(user)

}

// 通過接口來獲取任意參數，然後一一揭曉
func DoFiledAndMethod(any interface{}) {

	getType := reflect.TypeOf(any)
	fmt.Println("get Type is :", getType.Name())

	getValue := reflect.ValueOf(any)
	fmt.Println("get all Fields is:", getValue)

	// 獲取方法字段
	// 1. 先獲取interface的reflect.Type，然後通過NumField進行遍歷
	// 2. 再通過reflect.Type的Field獲取其Field
	// 3. 最後通過Field的Interface()得到對應的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 獲取方法
	// 1. 先獲取interface的reflect.Type，然後通過.NumMethod進行遍歷
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}