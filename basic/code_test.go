package basic

import (
	"fmt"
	"testing"
)

type Stu struct {
	Age int `json:"name"`
}

var m = map[string]Stu{
	"linx": Stu{18},
}

func Test1(t *testing.T) {
	//对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X， 但 go 中的 map 的 value 本身是不可寻址的
	//m["linx"].Age = 1

	//正确，对于不确定的key,比如外部传入的一定要判单
	if v, ok := m["linx"]; ok {
		v.Age = 1
	}
}

func Test2(t *testing.T) {
	i := make([]int, 5)
	i = append(i, 1, 2, 3)
	fmt.Println(i)

	j := make([]int, 0)
	j = append(j, 1, 2, 3, 4)
	fmt.Println(j)

	//	输出
	//	[0 0 0 0 0 1 2 3]
	//	[1 2 3 4]
	//	原因
	//	make如果输入值，会默认给其初始化默认值
}

func Test3(t *testing.T) {
	array := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range array {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Printf("key: %d, value: %d \n", k, *v)
	}

	//输出
	//key: 1, value: 3
	//key: 2, value: 3
	//key: 3, value: 3
	//key: 0, value: 3

	//解释
	//for range 循环的时候会创建每个元素的副本，而不是元素的引用，
	//所以 m[key] = &val 取的都是变量 val 的地址，所以最后map中的所有元素的值都是变量 val 的地址，
	//因为最后 val 被赋值为3，所有输出都是3.

	//for range还是不保证顺序的
}

func Test4(t *testing.T) {

	//map
	mapData := map[string]any{"age": 28, "name": "linx", "email": "1824517828@qq.com"}
	for key, val := range mapData {
		fmt.Printf("key=%v, val=%v\n", key, val)
	}

	//数组
	arr := []int{1, 2, 3, 4, 5, 6}
	for index, item := range arr {
		fmt.Printf("index=%v,item=%v\n", index, item)
	}

	//切片
	var slice []int = make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	slice = append(slice, 6)
	for index, item := range slice {
		fmt.Printf("index=%v,item=%v\n", index, item)
	}

	//map的循环是不保证顺序的
	//数组的循环是保证顺序的
	//切片的循环是保证顺序的
}

func Test5(t *testing.T) {

	//变量幽灵
	//短变量声明语法，很好用，但是代码块中使用短变更声明与外部相同的变量时，
	//没有语法编译错误，但是代码块中同名短变量声明从声明开始到代码块结束，对变量的修改将不会影响到外部变量！

	x := 1
	fmt.Println(x) //打印出来是1
	{
		fmt.Println(x) //打印出来1
		x = 3
		fmt.Println(x) //打印出来1

		x := 5
		fmt.Println(x) ////打印出来5

		x = 7
		fmt.Println(x) //打印出来7
	}
	fmt.Println(x) //打印出来3

	//打印结果
	//1
	//1
	//3
	//5
	//7
	//3
	//总结:代码块中定义的x无论怎么修改，都不会影响到外部
}
