package template

import "fmt"

type GetUp interface {
	// GetUp 起床
	GetUp(name string)
}

// WorkTemplate 每天工作搬砖模板
type WorkTemplate interface {
	//GetUp 起床
	GetUp

	// GoOut 出门
	GoOut(name string)

	// Ride 乘车
	Ride(name string)

	// Work1 工作
	Work1(name string)

	// GoHome 回家
	GoHome(name string)

	// Sleep 睡觉
	Sleep(name string)
}

// DoWork 模板方法
func DoWork(name string, work WorkTemplate) {
	work.GetUp(name)
	work.GoOut(name)
	work.Ride(name)
	work.Work1(name)
	work.GoHome(name)
	work.Sleep(name)
}

type Work struct {
}

func (Work) GetUp(name string) {
	fmt.Printf("%s---起床\n", name)
}
func (Work) GoOut(name string) {
	fmt.Printf("%s---出门\n", name)
}
func (Work) Ride(name string) {
	fmt.Printf("%s---乘车\n", name)
}
func (Work) Work1(name string) {
	fmt.Printf("%s---默认工作-->板砖\n", name)
}
func (Work) GoHome(name string) {
	fmt.Printf("%s---下班回家\n", name)
}
func (Work) Sleep(name string) {
	fmt.Printf("%s---睡觉\n", name)
}
