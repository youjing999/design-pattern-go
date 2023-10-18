package lazy_mode_native

import (
	"sync"
)

type Singleton struct {
	Name string
	Port int
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Name: "192.168.0.130",
			Port: 3306,
		}
	})
	return instance
}
