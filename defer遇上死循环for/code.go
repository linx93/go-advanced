package main

import (
	"fmt"
	"sync"
)

// 不会执行defer wg.Done(),会一直在for中循环阻塞，
func f1(wg *sync.WaitGroup) {
	defer wg.Done()

	//code ....
	fmt.Println("code...")

	for {
		//code ...
		fmt.Println(1)
	}
}

// 不会执行defer wg.Done(),panic直接退出
func f2(wg *sync.WaitGroup) {
	defer wg.Done()

	//code ....
	fmt.Println("code...")

	for {
		//code ...
		panic(123)

	}
}

// 不会执行defer wg.Done(),panic后进入recover，recover执行结束之后程序结束
func f3(wg *sync.WaitGroup) {
	defer wg.Done()

	r := recover()
	if r != nil {
		fmt.Println("recover:", r)
	}

	//code ....
	fmt.Println("code...")

	for {
		//code ...
		panic(123)

	}
}

// 会执行defer wg.Done(),panic后进入recover，recover执行结束之后程序继续
func f4(wg *sync.WaitGroup) {
	defer wg.Done()

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("recover:", r)
		}
	}()

	//code ....
	fmt.Println("code...")

	for {
		//code ...
		panic(123)

	}
}

var wg = sync.WaitGroup{}

func main() {
	//测试f1
	//wg.Add(1)
	//f1(&wg)
	//wg.Wait()
	//fmt.Println("结束main")

	//测试f2
	//wg.Add(1)
	//f2(&wg)
	//wg.Wait()
	//fmt.Println("结束main")

	//测试f3
	//wg.Add(1)
	//f3(&wg)
	//wg.Wait()
	//fmt.Println("结束main")

	//测试f4
	wg.Add(1)
	f4(&wg)
	wg.Wait()

	fmt.Println("结束main")
}
