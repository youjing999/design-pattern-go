package main

import "fmt"

// Iterator 泛型接口，表示一个迭代器
type Iterator[T any] interface {
	Next() T
	HasNext() bool
}

// SliceIterator 泛型结构体，实现 Iterator 接口用于切片迭代
type SliceIterator[T any] struct {
	data  []T
	index int
}

// NewSliceIterator 创建一个切片迭代器
func NewSliceIterator[T any](data []T) *SliceIterator[T] {
	return &SliceIterator[T]{data: data, index: -1}
}

// Next 返回下一个元素
func (it *SliceIterator[T]) Next() T {
	it.index++
	return it.data[it.index]
}

// HasNext 检查是否有下一个元素
func (it *SliceIterator[T]) HasNext() bool {
	return it.index+1 < len(it.data)
}

func main() {
	// 创建一个整数切片
	intSlice := []int{1, 2, 3, 4, 5}

	// 创建一个整数切片迭代器
	intIterator := NewSliceIterator(intSlice)

	// 使用迭代器遍历切片并打印元素
	for intIterator.HasNext() {
		fmt.Println(intIterator.Next())
	}

	// 创建一个字符串切片
	stringSlice := []string{"apple", "banana", "cherry"}

	// 创建一个字符串切片迭代器
	stringIterator := NewSliceIterator(stringSlice)

	// 使用迭代器遍历切片并打印元素
	for stringIterator.HasNext() {
		fmt.Println(stringIterator.Next())
	}
}
