package main

import (
	"fmt"
	"goDesignPattern/AbstractFactory"
	"goDesignPattern/Builder"
	"goDesignPattern/Dicorator"
	"goDesignPattern/Memento"
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

func testDicorator() {
	// 测试装饰者模式，用管道来管理。  注意这里，父类参数接受子类参数，只能通过接口，或者子接口（嵌套父接口）。 像我们这里，middleware 用子接口的方式来实现
	pipeline := Dicorator.Pipeline{}
	pipeline.Send(&Dicorator.Context{
		"请求经过",
	}).Through(&Dicorator.Middleware1{}, &Dicorator.Middleware2{}, &Dicorator.Middleware3{}).Then(&Dicorator.ConcreteBusinessComponent{})

	//Dicorator.TestAb(&(Dicorator.B{}))
}

func testMemento() {
	// 测试备忘录模式

	// 负责人
	Caretaker := &Memento.Caretaker{make([]*Memento.Memento, 0)}

	// 原发器
	Origintor := &Memento.Originator{}
	Origintor.SetState("one")
	fmt.Println("当前状态：" + Origintor.GetState())
	Caretaker.AddMemento(Origintor.CreateMemento())

	Origintor.SetState("Two")
	fmt.Println("当前状态：" + Origintor.GetState())
	Caretaker.AddMemento(Origintor.CreateMemento())

	Origintor.SetState("Three")
	fmt.Println("当前状态：" + Origintor.GetState())
	Caretaker.AddMemento(Origintor.CreateMemento())

	// 恢复最开始的
	Origintor.RestoreMemento(Caretaker.GetMemento(0))
	fmt.Println("恢复")
	fmt.Println("当前状态：" + Origintor.GetState())
}

func main() {
	testFactory()
	testBuilder()
	testComposite()
	testDicorator()
	testMemento()
}
