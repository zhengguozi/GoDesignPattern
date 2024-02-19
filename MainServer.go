package main

import (
	"goDesignPattern/AbstractFactory"
	"goDesignPattern/Builder"
	"goDesignPattern/Structure/Composite"
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

func testComposite() {
	// 测试结构型-组合模式

	file2 := &Composite.File{"File1-1"}
	file3 := &Composite.File{"File1-2"}
	fold1 := &Composite.Folder{Name: "File1-3"}

	file4 := &Composite.File{"File1-3-1"}
	file5 := &Composite.File{"File1-3-2"}
	fold1.AddComponent(file4)
	fold1.AddComponent(file5)

	fold2 := &Composite.Folder{Name: "Fold1"}
	fold2.AddComponent(file2)
	fold2.AddComponent(file3)
	fold2.AddComponent(fold1)
	fold2.Search("", "关关")
}

func main() {
	testFactory()
	testBuilder()
	testComposite()
}
