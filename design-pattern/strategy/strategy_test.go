package strategy

import (
	"fmt"
	"math"
	"testing"
)

func TestOperator_Calc(t *testing.T) {
	add := NewOperatorAdd()
	//设置加法策略
	calc, err := add.Calc(1, 1)
	if err != nil {
		fmt.Print("operator add failed: ", err)
		return
	}
	fmt.Printf("1+1=%v\n", calc)

	divide := NewOperatorDivide()
	//设置加法策略
	calc, err = divide.Calc(10, 2)
	if err != nil {
		fmt.Print("operator divide failed: ", err)
		return
	}
	fmt.Printf("10/2=%v\n", calc)

	sqrtSum := NewOperator(&SqrtSum{})
	calc, err = sqrtSum.Calc(2, 3)
	if err != nil {
		fmt.Print("operator divide failed: ", err)
		return
	}
	fmt.Printf("2^2 + 3^2 =%v\n", calc)
}

type SqrtSum struct{}

// 实现一个自己的策略
func (sq SqrtSum) Execute(a, b float64) (float64, error) {
	//a^2 + b^2
	return math.Pow(a, 2) + math.Pow(b, 2), nil
}
