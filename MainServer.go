package main

import "goDesignPattern/AbstractFactory"

func testFactory() {
	// 测试抽象工厂模式
	abstractFactory1 := AbstractFactory.FactoryHelper("huawei")
	AbstractFactory.UseFactory(abstractFactory1)
}

func main() {
	testFactory()
}
