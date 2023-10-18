package main

import (
	"sync"
	"sync/atomic"
)

type DbPool struct {
	Host string
	Port int
}

var dbPoolInit *DbPool
var lock sync.Locker

//goroutine1进来，实例化dbPoolInit到一半，goroutine2进来
// dbPoolInit读取到 != nil，返回未完全实例化的 dbPoolInit
var initialed uint32

//通过加锁解决并发
func GetDBPool_1() *DbPool {
	lock.Lock() // 如果实例存在没有必要加锁
	defer lock.Unlock()

	if dbPoolInit == nil {
		dbPoolInit = &DbPool{}
	}
	return dbPoolInit
}

func GetDBPool_2() *DbPool {

	//这边不是完全原子性
	if dbPoolInit == nil {
		lock.Lock()
		defer lock.Unlock()
		dbPoolInit = &DbPool{}
	}
	return dbPoolInit
}

func GetDbPool() *DbPool {

	if atomic.LoadUint32(&initialed) == 1 {
		return dbPoolInit
	}
	//加锁
	lock.Lock()
	defer lock.Unlock()

	if atomic.LoadUint32(&initialed) == 0 {
		dbPoolInit = &DbPool{}
		atomic.StoreUint32(&initialed, 1)
	}

	return dbPoolInit
}

//通过Once解决
var once sync.Once

func GetDbPoolByOnce() *DbPool {
	once.Do(func() {
		dbPoolInit = &DbPool{}
	})
	return dbPoolInit
}
