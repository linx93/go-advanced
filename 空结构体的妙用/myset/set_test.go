package myset

import (
	"fmt"
	"testing"
)

// 执行测试： go test -v set_test.go set.go  ,这里如果不带上set_test.go中引用的set.go文件，就会报错 undefined Set
func TestMySet_Add(t *testing.T) {
	set := Set[string]{}

	set.Add("e1")
	set.Add("e2")
	set.Add("e3")
	set.Add("e1")
	set.Add("e2")
	set.Add("e3")
	fmt.Println(set) //map[e1:{} e2:{} e3:{}]，可以看到打印出来的是满足set性质的

	exit1 := set.Contains("e1")
	exit2 := set.Contains("e4")
	fmt.Println(exit1) // true
	fmt.Println(exit2) // false

	set.Remove("e1")
	fmt.Println(set) //map[e2:{} e3:{}]

	set.String()

}
