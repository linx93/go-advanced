package main

import "fmt"

//有时候我们需要创建一组方法集的实现（一般来说是实现一个接口），
//但并不需要在这个实现中存储任何数据，这种情况下，我们可以使用空结构体来实现

type Linx interface {
	Work()
	PlayGame()
	PlayBasketball()
}

// X 空结构体作为方法接收器
type X struct{}

func (x *X) Work() {
	fmt.Println("Work")
}

func (x *X) PlayGame() {

}

func (x *X) PlayBasketball() {
	fmt.Println("PlayBasketball")
}
