package main

import (
	"fmt"
	"strings"
)

type DbConfig struct {
	Host string
	Port int
	User string
	Pwd  string
}

func NewDbConfig(Host string, Port int, User string, Pwd string) *DbConfig {
	return &DbConfig{
		Host: Host,
		Port: Port,
		User: User,
		Pwd:  Pwd,
	}
}

type DbBuilder struct {
	DbConfig
	err error
}

func Builder() *DbBuilder {
	b := new(DbBuilder)
	b.DbConfig.Host = "127.0.0.1"
	b.DbConfig.Port = 3306
	b.DbConfig.User = "root"
	b.DbConfig.Pwd = "root"

	return b
}
func (builder *DbBuilder) Host(Host string) *DbBuilder {

	if builder.err != nil {
		return builder
	}
	if strings.TrimSpace(Host) == "" {
		builder.err = fmt.Errorf("invalid Host is %s", Host)
	}
	builder.DbConfig.Host = Host
	return builder
}

func (builder *DbBuilder) Port(Port int) *DbBuilder {

	if builder.err != nil {
		return builder
	}
	if Port == 0 {
		builder.err = fmt.Errorf("invalid Port is %d", Port)
	}
	builder.DbConfig.Port = Port
	return builder
}

func (builder *DbBuilder) User(User string) *DbBuilder {

	if builder.err != nil {
		return builder
	}
	if strings.TrimSpace(User) == "" {
		builder.err = fmt.Errorf("invalid User is %s", User)
	}
	builder.DbConfig.User = User
	return builder
}

func (builder *DbBuilder) Pwd(Pwd string) *DbBuilder {

	if builder.err != nil {
		return builder
	}
	if strings.TrimSpace(Pwd) == "" {
		builder.err = fmt.Errorf("invalid Pwd is %s", Pwd)
	}
	builder.DbConfig.Pwd = Pwd
	return builder
}

func (builder *DbBuilder) Build() (*DbConfig, error) {

	if builder.err != nil {
		return nil, builder.err
	}

	return &builder.DbConfig, nil
}

func main() {
	build, err := Builder().Host("192.168.0.1").Port(3306).User("whisky").Pwd("xzq").Build()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(build)
}
