package iota的用法

//有个本质的理解，iota出现在第几行它的值就是几

//简单用法:

const (
	Stop = iota
	Run
)

const (
	Status1 = iota       //0
	Status2              //1
	Status3              //2
	Status4 = iota + 100 //3+100 ,注意iota出现在第几行它的值就是几
	Status5              //104
	Status6              //105
	Status7              //106
)

type ObjectType int

// 下面两种用法其实是一样的
const (
	ObjectString = ObjectType(iota)
	ObjectList
	ObjectSet
	ObjectZSet
	ObjectHash
)

const (
	ObjectString_ ObjectType = iota
	ObjectList_
	ObjectSet_
	ObjectZSet_
	ObjectHash_
)
