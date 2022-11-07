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
