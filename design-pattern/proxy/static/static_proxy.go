package static

import "fmt"

type Subject interface {
	Do(input string) string
}

type RealSubject struct{}

func (RealSubject) Do(input string) string {
	fmt.Printf("Do something\n")
	return fmt.Sprintf("[Do+%v+Do]", input)
}

type Proxy struct {
	realSubject RealSubject
}

func (proxy Proxy) Do(input string) string {
	// before 前置处理
	fmt.Printf("对%s做前置处理\n", input)
	output := proxy.realSubject.Do(input)
	//after 后置处理
	fmt.Printf("对%s做后置处理\n", input)
	return output
}
