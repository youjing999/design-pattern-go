package main

import "fmt"

/**
抽象方法
*/

//AbstractFactory 抽象的手机制造商，可以制造手机和手表
type AbstractFactory interface {
	CreateWatch() IWatch
	CreatePhone() ICallPhone
}

// IWatch 手表可以看时间
type IWatch interface {
	WatchTime()
}

// ICallPhone 手机可以打给某个人
type ICallPhone interface {
	CallSomebody()
}

//MIFactory 小米手机制造商可以制造手机和手表
type MIFactory struct {
}

func (mi *MIFactory) CreateWatch() IWatch {
	fmt.Println("制造小米手环")
	return &MIWatch{}
}
func (mi *MIFactory) CreatePhone() ICallPhone {
	fmt.Println("制造小米手机")
	return &MIPhone{}
}

//AppleFactory 苹果手机制造商可以制造手机和手表
type AppleFactory struct {
}

func (apple *AppleFactory) CreateWatch() IWatch {
	fmt.Println("制造Apple Watch")
	return &AppleWatch{}
}
func (apple *AppleFactory) CreatePhone() ICallPhone {
	fmt.Println("制造iPhone")
	return &IPhone{}
}

type MIWatch struct{}

func (miWatch *MIWatch) WatchTime() {
	fmt.Println("小米手环看时间")
}

type MIPhone struct{}

func (miPhone *MIPhone) CallSomebody() {
	fmt.Println("小米手机打电话")
}

type AppleWatch struct{}

func (watch *AppleWatch) WatchTime() {
	fmt.Println("Apple Watch看时间")
}

type IPhone struct{}

func (phone *IPhone) CallSomebody() {
	fmt.Println("iPhone打电话给某人")
}

func main() {
	var factory AbstractFactory
	var watch IWatch
	var phone ICallPhone

	factory = &MIFactory{}
	watch = factory.CreateWatch()
	watch.WatchTime()
	phone = factory.CreatePhone()
	phone.CallSomebody()

	fmt.Println("------------")
	factory = &AppleFactory{}
	appleWatch := factory.CreateWatch()
	appleWatch.WatchTime()
	iPhone := factory.CreatePhone()
	iPhone.CallSomebody()
}
