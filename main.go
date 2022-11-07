package main

import (
	"fmt"
	"log"
)

func main() {
	tiger := Tiger{}
	tiger.Run("[跑起来]")

	var swallow Bird = Swallow{}
	swallow.Run("[跑起来]")
	swallow.Fly("[飞起来]")
	var m = map[string]Stu{
		"linx": Stu{18},
	}

	if v, ok := m["linx"]; ok {
		v.Age = 1
	}

	//
	linx := Stu{Age: 18}
	testStr := "unknown"
	testStr1 := "unknown1"
	defer func(testStr1 string, stu Stu) {
		log.Printf("testStr=%v", testStr)
		log.Printf("testStr1=%v", testStr1)
		log.Printf("stu=%#v", linx)
		log.Printf("stu1=%#v", stu)
	}(testStr1, linx)
	testStr = fmt.Sprintf("德玛西亚")
	testStr1 = fmt.Sprintf("德玛西亚1")
	linx.Age = 28

}

type Stu struct {
	Age int `json:"age"`
}

// Animal 动物
type Animal interface {
	Run(string) string
}

// Bird 鸟 继承了 Animal 动物
type Bird interface {
	Animal
	Fly(string) string
}

// Tiger 老虎
type Tiger struct {
}

func (t Tiger) Run(str string) string {
	log.Printf("老虎run-->%s", str)
	return fmt.Sprintf("老虎run-->%s", str)
}

// Swallow 燕子
type Swallow struct {
}

func (t Swallow) Fly(str string) string {
	log.Printf("燕子Fly-->%s", str)
	return fmt.Sprintf("燕子Fly-->%s", str)
}

func (t Swallow) Run(str string) string {
	log.Printf("燕子Fly-->%s", str)
	return fmt.Sprintf("燕子Fly-->%s", str)
}
