package main

//
///**
//使用go泛型实现迭代器模式
//*/
//
//// Iterator 迭代器
//type Iterator[T any] interface {
//	HashNext() bool
//	Next() T
//}
//
//type MyIterator[T any] struct {
//	data  []*T
//	index int
//}
//
//func NewMyIterator[T any](data []*T) *MyIterator {
//	return &MyIterator{
//		data:  data,
//		index: 0,
//	}
//}
//
//func (my *MyIterator) HashNext() bool {
//	return my.index < len(my.data)
//}
//
//func (my *MyIterator) Next() *T {
//	item := my.data[my.index]
//	my.index++
//	return item
//}
