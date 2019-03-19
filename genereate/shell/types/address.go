package types

import "fmt"

type Address struct {
	Street string
	Town string
}

func (a *Address)String() string {
	return fmt.Sprintf("%v\n%v", a.Street, a.Town)
}

type AddressList []*Address

type AddressToBool func(*Address) bool

func (al AddressList)Filter(f AddressToBool) AddressList {
	var ret AddressList
	for _, a := range al {
		if f(a) {
			ret = append(ret, a)
		}
	}
	return ret
}