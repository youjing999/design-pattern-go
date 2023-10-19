package main

import "fmt"

type Greeter interface {
	Greet(name string)
}

type SimpleGreeter struct{}

func (s *SimpleGreeter) Greet(name string) {
	fmt.Printf("%s, hello\n", name)
}

type DecoratorGreeter struct {
	Greeter
}

func (d *DecoratorGreeter) Greet(name string) {
	//前置操作
	fmt.Println("装饰器前置")
	d.Greeter.Greet(name)
	//后置操作
	fmt.Println("装饰器后置")
}

func main() {
	greeter := &SimpleGreeter{}
	greeter.Greet("yj")

	decoratorGreeter := DecoratorGreeter{
		Greeter: greeter,
	}

	decoratorGreeter.Greet("whisky")
}
