package main

import "fmt"

//Target 目标接口，最后需要使用request方法
type Target interface {
	request() string
}

//Adapted 被适配的接口
type Adapted interface {
	translateRequest() string
}

func NewAdapted() Adapted {
	return &AdaptedImpl{}
}

// AdaptedImpl Adapted的一个实现，提供translateRequest具体实现
type AdaptedImpl struct {
}

func (a *AdaptedImpl) translateRequest() string {
	return "adapted method:translateRequest()"
}

type Adapter struct {
	Adapted
}

func NewAdapter(adapted Adapted) *Adapter {
	return &Adapter{
		Adapted: adapted,
	}
}

func (adapter *Adapter) request() string {
	return adapter.translateRequest()
}

func main() {
	adapted := NewAdapted()
	adapter := NewAdapter(adapted)
	request := adapter.request()
	fmt.Println(request)
}
