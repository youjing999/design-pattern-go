package main

import "fmt"

// Observer 观察者接口
type Observer interface {
	Update(string)
}

// Subject 主题接口
type Subject interface {
	Register(Observer)
	Remove(Observer)
	Notify(string)
}

//ConObserver 观察者实现
type ConObserver struct {
	name string
}

func (c *ConObserver) Update(msg string) {
	fmt.Printf("%s 接收到消息: %s\n", c.name, msg)
}

//ConSubject 具体主题
type ConSubject struct {
	observers []Observer
}

func (c *ConSubject) Register(observer Observer) {
	c.observers = append(c.observers, observer)
}

func (c *ConSubject) Remove(observer Observer) {
	for i, o := range c.observers {
		if o == observer {
			c.observers = append(c.observers[:i], c.observers[i+1:]...)
			break
		}
	}
}

func (c *ConSubject) Notify(msg string) {
	for _, observer := range c.observers {
		observer.Update(msg)
	}
}

func main() {
	subject := &ConSubject{}

	observer1 := &ConObserver{name: "Observer 1"}
	observer2 := &ConObserver{name: "Observer 2"}
	observer3 := &ConObserver{name: "Observer 3"}

	subject.Register(observer1)
	subject.Register(observer2)
	subject.Register(observer3)

	subject.Notify("Hello, observers!")
	subject.Remove(observer2)
	subject.Notify("Observer 2 has been unregistered.")
}
