package lazy_mode

import (
	"sync"
	"sync/atomic"
)

type Singleton struct {
	Host string
	Port int
}

var initialized uint32
var instance *Singleton
var mu sync.Locker

func GetInstance() *Singleton {

	if atomic.LoadUint32(&initialized) == 1 { // 原子操作
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		instance = &Singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
