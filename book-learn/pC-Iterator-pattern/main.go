package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type MyIterator struct {
	slice []interface{}
	index int
}

func NewMyIterator(slice []interface{}) *MyIterator {
	return &MyIterator{
		slice: slice,
		index: 0,
	}
}

func (my *MyIterator) HasNext() bool {
	return my.index < len(my.slice)
}

func (my *MyIterator) Next() interface{} {
	if my.HasNext() {
		val := my.slice[my.index]
		my.index++
		return val
	}
	return nil
}

func main() {
	slice := []interface{}{1, 2, 3, 4, 5}
	iterator := NewMyIterator(slice)
	for iterator.HasNext() {
		fmt.Println(iterator.Next())
	}
}
