package main

import (
	"fmt"
	"github.com/linx93/go-advanced/option_init_instance/init_code"
)

func main() {
	ins := init_code.New()
	fmt.Println(ins)

	ins1 := init_code.New(init_code.WithHttpProxyConfig("./config.yaml"),
		init_code.WithName("linx"),
		init_code.WithVersion("1.0.0"))
	fmt.Println(ins1)

	//个人感觉没有java或者kotlin中的方式体验好

}
