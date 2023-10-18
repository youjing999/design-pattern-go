package main

/**
简单工厂
*/
import "fmt"

type Printer interface {
	Print(name string) string
}

// CnPrinter Chinese
type CnPrinter struct {
	name string
}

func (cn *CnPrinter) Print(name string) string {

	return "中文打印机" + name
}

// EnPrinter English
type EnPrinter struct {
	name string
}

func (en *EnPrinter) Print(name string) string {

	return "英文打印机" + name
}

// NewPrinter 简单工厂
func NewPrinter(pName string) Printer {
	switch pName {
	case "cn":
		return new(CnPrinter)
	case "en":
		return new(EnPrinter)
	default:
		return new(CnPrinter)
	}
}

func main() {
	printer := NewPrinter("en")
	fmt.Println(printer.Print("willy"))
}
