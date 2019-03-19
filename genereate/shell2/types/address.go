package types

import "fmt"

//go:generate ./newList.sh Address
type Address struct {
	Street string
	Town string
}

func (a *Address)String() string {
	return fmt.Sprintf("%v\n%v", a.Street, a.Town)
}