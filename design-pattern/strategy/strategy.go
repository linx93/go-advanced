package strategy

import (
	"errors"
	"math"
)

// IStrategy 策略接口
type IStrategy interface {
	Execute(float64, float64) (float64, error)
}

// add 加
type add struct {
}

func (ad *add) Execute(a, b float64) (float64, error) {
	return a + b, nil
}

// sub 减
type sub struct {
}

func (sb *sub) Execute(a, b float64) (float64, error) {
	return a - b, nil
}

// multi 乘
type multi struct {
}

func (mi *multi) Execute(a, b float64) (float64, error) {
	return a * b, nil
}

// divide 除
type divide struct {
}

func (de *divide) Execute(a, b float64) (float64, error) {
	if b == 0 {
		return math.NaN(), errors.New("by zero error")
	}
	return a / b, nil
}

// operator 具体策略的执行者
type operator struct {
	strategy IStrategy
}

func NewOperator(strategy IStrategy) *operator {
	return &operator{strategy: strategy}
}

func NewOperatorAdd() *operator {
	return NewOperator(&add{})
}

func NewOperatorMulti() *operator {
	return NewOperator(&multi{})
}

func NewOperatorSub() *operator {
	return NewOperator(&sub{})
}

func NewOperatorDivide() *operator {
	return NewOperator(&divide{})
}

// SetStrategy 设置策略
func (operator *operator) SetStrategy(strategy IStrategy) {
	operator.strategy = strategy
}

// Calc 调用策略中的方法做计算
func (operator *operator) Calc(a, b float64) (float64, error) {
	return operator.strategy.Execute(a, b)
}
