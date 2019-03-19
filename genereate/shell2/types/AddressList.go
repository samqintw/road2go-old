package types

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
