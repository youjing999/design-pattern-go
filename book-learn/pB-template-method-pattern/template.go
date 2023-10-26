package main

import "fmt"

// AbstractClass 抽象类
type AbstractClass interface {
	TemplateMethod()
	AbstractMethod()
	HookMethod()
}

// ConcreteClass 具体子类
type ConcreteClass struct{}

func (c *ConcreteClass) TemplateMethod() {
	fmt.Println("TemplateMethod: Start")

	c.AbstractMethod()
	c.HookMethod()

	fmt.Println("TemplateMethod: End")
}

// AbstractMethod 抽象方法
func (c *ConcreteClass) AbstractMethod() {
	fmt.Println("AbstractMethod: Implemented by ConcreteClass")
}

// HookMethod 钩子方法
func (c *ConcreteClass) HookMethod() {
	fmt.Println("HookMethod: Default implementation in ConcreteClass")
}

func main() {
	abstractClass := &ConcreteClass{}

	abstractClass.TemplateMethod()
}
