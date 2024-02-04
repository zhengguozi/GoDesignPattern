package main

import (
	"goDesignPattern/AbstractFactory"
	"goDesignPattern/Builder"
)

func testFactory() {
	// 测试抽象工厂模式
	abstractFactory1 := AbstractFactory.FactoryHelper("huawei")
	AbstractFactory.UseFactory(abstractFactory1)
}

func testBuilder() {
	// 测试生成器模式
	director := Builder.NewDiretor()
	director.ShowCar(director.MakeCar())

	// 换个生成器
	director.SetBuilder(
		Builder.WenjieBuilder{
			Builder.BaseBuilder{},
		})

	director.ShowCar(director.MakeCar())
}

func main() {
	testFactory()
	testBuilder()
}
