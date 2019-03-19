package main
import (
"fmt"
)
type Animal interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

//1
func (c *Cat) Speak() string {
	return "Meow!"
}
type Llama struct {
}
func (l Llama) Speak() string {
	return "?????"
}
type JavaProgrammer struct {
}
func (j JavaProgrammer) Speak() string {
	return "Design patterns!"
}

// the cat's receiver is passing by point that is different from the others.
func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
