package Strategy

import "fmt"

// CreditCardStrategy 是具体的策略
type CreditCardStrategy struct {
	CardNumber string
	CVV        string
}

// Pay 实现 PaymentStrategy 接口
func (c *CreditCardStrategy) Pay(amount float64) string {
	fmt.Printf("Paid %.2f using Credit Card (%s)", amount, c.CardNumber)
	return ""
}

// 在文件初始化时自动注册策略
func init() {
	fmt.Println("Registering CreditCardStrategy")
	GlobalPaymentStrategyFactory.RegisterStrategy("credit_card", &CreditCardStrategy{
		CardNumber: "1234-5678-9876-5432",
		CVV:        "123",
	})
}
