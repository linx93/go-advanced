package main

import "fmt"

func main() {
	//在go中，每个指针变量都有两个与之相关的值，<type,value>,事实上每个变量都必须指定
	//一个类型，这就是为什么不能给未定义的变量赋空值，比如x:=nil,因为不知道x是什么类型的

	//例1.  nil!=nil的例子
	var a *string = nil               // a is <*string,nil>
	var b interface{} = nil           // b is <nil,nil>
	fmt.Println("a==nil :", a == nil) //true,
	fmt.Println("b==nil :", b == nil) //true,
	fmt.Println("a==b :", a == b)     //false,这里两个值为nil的变量不相等

	//因为:当a==nil判断时，==运算符比较类型和值, 当nil和对象比较时候，nil的类型和与它比较的对象声明的类型相同，
	//所以:
	//a==nil会被认为是<*string,nil>==<*string,nil>,所以返回true
	//b==nil会被认为是<nil,nil>==<nil,nil>,所以返回true
	//a==b会被认为是<*string,nil>==<nil,nil>,所以返回false

	//例2.  变量的声明类型和实际类型不是一个概念，下面的bb变量被赋值为aa之后，变量bb的声明类型对应为<nil,nil>而实际类型为<*string,nil>
	var aa *string = nil    //aa is <*string,nil>
	var bb interface{} = aa // bb is <*string,nil> 注意:var b interface{} =a赋值操作，并不会改变 b变量的声明类型。

	fmt.Println("aa == nil:", aa == nil) //true (<*int, nil>==<*int, nil>)
	fmt.Println("bb == nil:", bb == nil) //false (<*string,nil>==<nil,nil>
	fmt.Println("aa == bb:", aa == bb)   // true (<*string,nil>==<*string,nil>)

	//interface在go中比较比较特殊，它的类型是在运行时确定的
	//例3. interface类型说明

	var aa1 int = 12

	var bb1 float64 = 12

	var cc1 interface{} = aa1

	fmt.Println("aa1==12:", aa1 == 12) // true

	fmt.Println("bb1==12:", bb1 == 12) // true

	fmt.Println("cc1==12:", cc1 == 12) // true

	fmt.Println("aa1==cc1:", aa1 == cc1) // true

	fmt.Println("bb1==cc1:", bb1 == cc1) // false <float64,12>和<int,12>比肯定是false

	fmt.Printf("aa1 <%T,%v>\n", aa1, aa1) // aa1 <int,12>

	fmt.Printf("bb1 <%T,%v>\n", bb1, bb1) // bb1 <float64,12>

	fmt.Printf("cc1 <%T,%v>\n", cc1, cc1) // cc1 <int,12>

	cc1 = bb1

	fmt.Printf("cc1 <%T,%v>\n", cc1, cc1) // cc1 <float64,12> 这里再次赋值，导致cc1的类型再次发生改变

}
