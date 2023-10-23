package main

import "fmt"

/**
享元模式
*/

// Flyweight 结构体表示享元对象，其中包含一个共享的数据 sharedData。
// Operation 方法用于执行操作，它接收一个唯一的数据 uniqueData，并将共享数据和唯一数据打印出来。
type Flyweight struct {
	sharedData string
}

func (f *Flyweight) Operation(uniqueData string) {
	fmt.Printf("Shared data: %s, Unique data: %s\n", f.sharedData, uniqueData)
}

//FlyweightFactory 定义享元工厂
type FlyweightFactory struct {
	flyweights map[string]*Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	flyweightsMap := make(map[string]*Flyweight)
	flyweightsMap["sharedData1"] = &Flyweight{
		"go sharedData1",
	}

	flyweightsMap["sharedData2"] = &Flyweight{
		"go sharedData2",
	}
	return &FlyweightFactory{
		flyweights: make(map[string]*Flyweight),
	}

}

func (ff *FlyweightFactory) GetFlyweight(key string) *Flyweight {
	if flyweight, ok := ff.flyweights[key]; ok {
		return flyweight
	}

	flyweight := &Flyweight{sharedData: key}
	ff.flyweights[key] = flyweight
	return flyweight
}

func main() {
	factory := NewFlyweightFactory()

	flyweight1 := factory.GetFlyweight("sharedData1")
	flyweight1.Operation("uniqueData1")

	flyweight2 := factory.GetFlyweight("sharedData2")
	flyweight2.Operation("uniqueData2")
}
