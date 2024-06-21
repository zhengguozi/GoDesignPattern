package Strategy

import "fmt"

// PaymentStrategy 支付策略接口
type PaymentStrategy interface {
	Pay(amount float64) string
}

// PaymentStrategyFactory 生产PaymentStrategy的工厂
type PaymentStrategyFactory struct {
	strategies map[string]PaymentStrategy
}

// GlobalStrategyFactory 是全局策略工厂实例
var GlobalPaymentStrategyFactory = NewPaymentStrategyFactory()

// NewPaymentStrategyFactory 创建一个新的 PaymentStrategyFactory
func NewPaymentStrategyFactory() *PaymentStrategyFactory {
	fmt.Println("Initializing GlobalStrategyFactory")
	return &PaymentStrategyFactory{
		strategies: make(map[string]PaymentStrategy),
	}
}

// RegisterStrategy 注册一个策略
func (f *PaymentStrategyFactory) RegisterStrategy(name string, strategy PaymentStrategy) {
	f.strategies[name] = strategy
}

// GetStrategy 返回一个策略
func (f *PaymentStrategyFactory) GetStrategy(name string) (PaymentStrategy, bool) {
	strategy, found := f.strategies[name]
	return strategy, found
}
