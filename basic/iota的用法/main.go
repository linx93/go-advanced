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

const (
	status1 = iota       //  0 注意iota出现在第几行它的值就是几
	status2 = iota - 1   //  0
	status3 = iota - 1   //  1
	status4 = iota + 100 //  103
	status5 = iota - 1   //  3
	status6 = iota - 1   //  4
	status7 = iota - 1   //  5
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
