package basic

import (
	"fmt"
	"testing"
)

// golang中几种初始化的区别
// 1. new(type)的方式，返回的是type的指针类型，会默认初始化结构体的数据，其实就和java的无参构造类似，
// 2 (*Type)(nil)的方式，返回的是type的指针类型，不会初始化结构体的数据
// 3. type{}的方式，返回的是结构体类型，会初始化数据
func TestTypeCheck(t *testing.T) {
	//var v any = new(Bird)
	//这个情况打印结果:
	//0xc00005c610 Bird{Name=,Color=,Age=0} 这里的new可以理解为地址存在，同时也初始化对象，其实和java中的new一个道理
	//Animal type[*basic.Bird] Bird{Name=,Color=,Age=0} 类型是*basic.Bird，

	//var v any = (*Dog)(nil)
	//这个情况打印结果:
	//0xc00000a038 <nil>  可以理解为地址存在，但是没有初始化对象
	//*Dog type[*basic.Dog] <nil> 类型是*basic.Dog，但是没有初始化数据

	//var v any
	//这个情况打印结果:
	//0xc00005c610 <nil> 没初始化数据
	//nil type[<nil>] <nil>  没类型且没数据

	var v any = new(Dog)
	//这个情况打印结果:
	//0xc00000a038 Dog{Name=,Age=0}  new就是开辟地址且初始化对象数据，和java的new类似，相当于golang中只有一个无参构造
	//*Dog type[*basic.Dog] Dog{Name=,Age=0}

	//var v any = Dog{}
	//这个情况打印结果:
	//0xc0001180d8 Dog{Name=,Age=0} 这里初始化地址且初始化对象数据，不同于new的是这里的类型是结构体而不是结构体的指针
	//Dog type[basic.Dog] Dog{Name=,Age=0} 类型是结构体而不是指针

	switch v := v.(type) {
	case Animal:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Animal type[%T] %v\n", v, v)
	case *Animal:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Animal type[%T] %v\n", v, v)
	case Dog:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("Dog type[%T] %v\n", v, v)
	case *Dog:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("*Dog type[%T] %v\n", v, v)
	case nil:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("nil type[%T] %v\n", v, v)
	default:
		fmt.Printf("%p %v\n", &v, v)
		fmt.Printf("unknow type\n")
	}

	//var animal1 Animal = (*Dog)(nil)  编译通不过，因为dog并没有实现animal
	var animal Animal = (*Bird)(nil)
	fmt.Printf("%p %v\n", &animal, animal)
	fmt.Printf("type[%T] %v\n", animal, animal)
	//这个情况打印结果:
	//0xc0000885f0 <nil>  没初始化
	//type[*basic.Bird] <nil> 类型是指针类型*basic.Bird，没初始胡数据
}

type Animal interface {
	GetName() string
	GetAge() uint8
	GetColor() string
}

type Bird struct {
	Name  string
	Color string
	Age   uint8
}

type Dog struct {
	Name string
	Age  uint8
}

// bird implement Animal
func (b *Bird) GetName() string {
	return b.Name
}

func (b *Bird) GetColor() string {
	return b.Color
}

func (b *Bird) GetAge() uint8 {
	return b.Age
}

func (b Bird) String() string {
	return fmt.Sprintf("Bird{Name=%v,Color=%v,Age=%v}", b.Name, b.Color, b.Age)
}

// dog not implement Animal
func (d *Dog) GetName() string {
	return d.Name
}

func (d Dog) String() string {
	return fmt.Sprintf("Dog{Name=%v,Age=%v}", d.Name, d.Age)
}
