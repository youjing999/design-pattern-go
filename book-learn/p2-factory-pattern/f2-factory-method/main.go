package main

import "fmt"

/**
工厂方法
*/

//CalculateHandle 计算器（产品）接口
//每个计算器都需要输入两个数,计算结果
type CalculateHandle interface {
	setNum1(int)
	setNum2(int)
	calResult() int
}

//BaseCalculate 每个计算操作都需要两个数，将这两个数抽出来
type BaseCalculate struct {
	num1 int
	num2 int
}

//PlusCalculate 加法
type PlusCalculate struct {
	*BaseCalculate
}

// MinCalculate 减法
type MinCalculate struct {
	*BaseCalculate
}

//给num1 num2赋值
func (baseCalculate *BaseCalculate) setNum1(num int) {
	baseCalculate.num1 = num
}

func (baseCalculate *BaseCalculate) setNum2(num int) {
	baseCalculate.num2 = num
}

func (p *PlusCalculate) calResult() int {
	return p.num1 + p.num2
}

func (m *MinCalculate) calResult() int {
	return m.num1 - m.num2
}

//CalculateFactory 计算器工厂生产计算器
type CalculateFactory interface {
	Create() CalculateHandle
}

//PlusFactory 加法工厂结构体
type PlusFactory struct {
}

func (p *PlusFactory) Create() CalculateHandle {
	return &PlusCalculate{
		BaseCalculate: &BaseCalculate{},
	}
}

//MinFactory 减法工厂结构体
type MinFactory struct {
}

func (m *MinFactory) Create() CalculateHandle {
	return &MinCalculate{
		BaseCalculate: &BaseCalculate{},
	}
}

func main() {
	var factory CalculateFactory
	factory = &PlusFactory{}
	plusOp := factory.Create()
	plusOp.setNum1(1)
	plusOp.setNum2(4)
	fmt.Printf("加法计算结果:%d\n", plusOp.calResult())

	factory = &MinFactory{}
	MinOp := factory.Create()
	MinOp.setNum1(99)
	MinOp.setNum2(1)
	fmt.Printf("减法计算结果:%d\n", MinOp.calResult())
}
