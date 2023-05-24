package myset

import (
	"fmt"
	"strings"
)

type Set[E comparable] map[E]struct{}

// Add 添加元素
func (set Set[E]) Add(element E) {
	//元素放到map的key，value用空结构体填充
	set[element] = struct{}{}
}

// Remove 移除元素
func (set Set[E]) Remove(element E) {
	//本质是移除map的key
	delete(set, element)
}

// Contains 存在元素
func (set Set[E]) Contains(element E) bool {
	//本质就是获取map中的k-v
	_, ok := set[element]
	return ok
}

func (set Set[E]) String() {
	s := "["
	for k := range set {
		s += fmt.Sprintf("%v,", k)
	}
	trim := strings.Trim(s, ",")
	trim += "]"
	fmt.Println(trim)
}
