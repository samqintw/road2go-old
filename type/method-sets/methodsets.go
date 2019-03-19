package main

import "fmt"

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

//Note that both pointer and value methods can both be called on both pointer and non-pointer values.
func main() {
	// A bare value
	var lst List
	// lst have only method Len()
	lst.Append(1)
	// implicitly translates to
	// (&lst).Append(1)
	fmt.Printf("%v (len: %d)\n", lst, lst.Len())

	// INVALID: Append has a pointer receiver
	// CountInto(lst, 1, 10)
	CountInto(&lst, 1, 10)
	if LongEnough(lst) {  // VALID: Identical receiver type
		fmt.Printf(" - lst is long enough")
	}

	// A pointer value
	pList := new(List)
	pList.Append(2)
	// implicitly translates to
	// (*pList).Len()
	fmt.Printf("%v (len: %d)\n", pList, pList.Len())
	// VALID: Identical receiver type
	CountInto(pList, 1, 10)
	if LongEnough(pList) {  // VALID: a *List can be dereferenced for the receiver
		fmt.Printf(" - plst is long enough")
	}

	//================
	//lists1 := map[string]List{}
	//cannot be rewritten as (&lists["primes"]).Append(7)
	//lists1["primes"].Append(7)

	lists2 := map[string]*List{}
	lists2["primes"] = new(List)
	lists2["primes"].Append(7)
	count := lists2["primes"].Len() // can be rewritten as (*lists["primes"]).Len()
	fmt.Println(count)


}