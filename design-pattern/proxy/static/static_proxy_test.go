package static

import (
	"fmt"
	"testing"
)

func TestProxy_Do(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do("i am linx")
	fmt.Println("结果：", res)
}
