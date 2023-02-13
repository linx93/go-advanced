package main

import (
	"fmt"
	"github.com/linx93/go-advanced/option_init_instance/init_code"
)

func main() {
	//1. 参考grpc中创建方式
	ins := init_code.New()
	fmt.Println(ins)

	ins1 := init_code.New(init_code.WithHttpProxyConfig("./config.yaml"),
		init_code.WithName("linx"),
		init_code.WithVersion("1.0.0"))
	fmt.Println(ins1)

	//2. 参考java中的方式
	i := init_code.NewIns().SetAge(18).SetName("linx").SetConfig("./config.yaml")
	fmt.Println(i)

}
