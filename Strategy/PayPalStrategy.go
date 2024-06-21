package Strategy

import "fmt"

// PayPalStrategy 是具体的策略
type PayPalStrategy struct {
	Email string
}

// Pay 实现 PaymentStrategy 接口
func (p *PayPalStrategy) Pay(amount float64) string {
	fmt.Printf("Paid %.2f using PayPal (%s)", amount, p.Email)
	return ""
}

// 在文件初始化时自动注册策略
func init() {
	fmt.Println("Registering PayPalStrategy")
	GlobalPaymentStrategyFactory.RegisterStrategy("paypal", &PayPalStrategy{
		Email: "user@example.com",
	})
}
