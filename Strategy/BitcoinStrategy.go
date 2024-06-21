package Strategy

import "fmt"

// BitcoinStrategy 是具体的策略
type BitcoinStrategy struct {
	WalletAddress string
}

// Pay 实现 PaymentStrategy 接口
func (b *BitcoinStrategy) Pay(amount float64) string {
	fmt.Printf("Paid %.2f using Bitcoin (%s)", amount, b.WalletAddress)
	return ""
}

// 在文件初始化时自动注册策略
func init() {
	fmt.Println("Registering BitcoinStrategy")
	GlobalPaymentStrategyFactory.RegisterStrategy("bitcoin", &BitcoinStrategy{
		WalletAddress: "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	})
}
