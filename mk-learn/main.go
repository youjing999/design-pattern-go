package main

import "fmt"

type DbOptions struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
}

type Option func(options *DbOptions)

// 这个函数主要用来设置Host

func WithHost(host string) Option {
	return func(options *DbOptions) {
		options.Host = host
	}
}

func NewOpts(options ...Option) *DbOptions {
	//先实例化好dbOptions，填充上默认值
	dbOpts := &DbOptions{
		Host: "127.0.0.1",
		Port: 3306,
	}

	for _, option := range options {
		option(dbOpts)
	}
	return dbOpts
}

func main() {
	//opt := NewOpts()
	opt := NewOpts(WithHost("192.168..46.130"))

	fmt.Println(opt)
}
