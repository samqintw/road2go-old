// https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
package main

import (
	"fmt"
)

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

// f1 = f2
func f2() (result int) {
	result = 0  //return语句不是一条原子调用，return xxx 其实是赋值 ＋ret指令
	func() {    //defer被插入到return之前执行，也就是 赋返回值 和 ret指令 之间
		result++
	}()
	return
}

// f3 = f4
func f3() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f4() (r int) {
	t := 5
	r = t
	func() {
		t = t + 5
	}()
	return
}

// f5 = f6
func f5() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func f6() (r int) {
	r = 1  //给返回值赋值
	func(r int) {        //这里改的r是传值传进去的r，不会改变要返回的那个r值
		r = r + 5
	}(r)
	return        //空的return
}

func main()  {
	fmt.Println("f1() :", f1())
	fmt.Println("f2() :", f2())
	fmt.Println("f3() :", f3())
	fmt.Println("f4() :", f4())
	fmt.Println("f5() :", f5())
	fmt.Println("f6() :", f6())
}
