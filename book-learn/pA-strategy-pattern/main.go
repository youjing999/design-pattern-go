package main

import "fmt"

/**
策略模式
*/

type PaymentStrategy interface {
	Pay(money float64) string
}

type WxPay struct {
}

func (wx *WxPay) Pay(money float64) string {
	return fmt.Sprintf("使用 wx支付了 %.2f", money)
}

type ZfbPay struct {
}

func (zfb *ZfbPay) Pay(money float64) string {
	return fmt.Sprintf("使用 zfb支付了 %.2f", money)
}

type ContextStrategy struct {
	paymentStrategy PaymentStrategy
}

func (c *ContextStrategy) SetStrategy(paymentStrategy PaymentStrategy) {
	c.paymentStrategy = paymentStrategy
}

func (c *ContextStrategy) DoTask(money float64) string {
	return c.paymentStrategy.Pay(money)
}

func main() {
	ctx := &ContextStrategy{}
	ctx.SetStrategy(&WxPay{})
	task := ctx.DoTask(200)
	fmt.Println(task)

	fmt.Println("---------")
	ctx.SetStrategy(&ZfbPay{})
	doTask := ctx.DoTask(300)
	fmt.Println(doTask)
}
