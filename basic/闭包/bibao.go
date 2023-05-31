package main

import "fmt"

//闭包概念:函数+函数引用的环境

func main() {
	//实例1
	//调用f1拿到一个返回，此返回是一个func
	f := f1()
	//调用f
	v1 := f(1)
	fmt.Println(v1) //0+1=1

	v2 := f(1)
	fmt.Println(v2) //1+1=2

	v3 := f(1)
	fmt.Println(v3) //2+1=3

	//实例2
	B() //2
	C() //2
}

// 实例1 匿名行数+引用的环境是变量initVal
func f1() func(i int) int {
	var initVal = 0
	return func(i int) int {
		initVal = initVal + 1
		//这里返回initVal变量，导致了变量逃逸到堆上
		return initVal
	}
}

// 实例2
func A(i int) {
	i++
	fmt.Println(i)
}

func B() {
	fb := A
	fb(1)
}

func C() {
	fc := A
	fc(1)
}
