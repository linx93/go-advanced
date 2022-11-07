package template

import (
	"fmt"
	"testing"
)

func TestWork_DoWork(t *testing.T) {
	work := Work{}
	DoWork("linx", work)

	fmt.Println("---------------------------------------")

	coderWork := CoderWork{}
	DoWork("linx", coderWork)
}

// 开发者的工作
type CoderWork struct {
	Work
}

func (CoderWork) Work1(name string) {
	//实现coder的工作内容
	fmt.Printf("%s---工作-->写java、golang、kotlin等等等\n", name)
}
