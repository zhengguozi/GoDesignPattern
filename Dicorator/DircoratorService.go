package Dicorator

import "fmt"

type Context struct {
	RequestName string
}

type Component interface {
	Handle(context *Context)
}

// 具体业务逻辑实现类
type ConcreteBusinessComponent struct {
}

func (c *ConcreteBusinessComponent) Handle(context *Context) {
	fmt.Println("我是业务逻辑处理")
	fmt.Println("最终收到的参数名为：" + context.RequestName)
}

// 中间件(子接口)，其实就是装饰者，本身也实现了Component，只不过多了一个可以设置其他component的方法
type Middleware interface {
	Component
	setComponent(component Component)
}

type BaseMiddleware struct {
	component Component
}

func (m *BaseMiddleware) setComponent(component Component) {
	m.component = component
}

func (m *BaseMiddleware) Handle(context *Context) {
	fmt.Println("父中间件打印")
}

type Middleware1 struct {
	BaseMiddleware
}

func (m *Middleware1) Handle(context *Context) {
	fmt.Println("中件间1前置处理")
	context.RequestName = context.RequestName + "中间件1-"
	m.component.Handle(context)
	fmt.Println("中件间1后置处理")
}

type Middleware2 struct {
	BaseMiddleware
}

func (m *Middleware2) Handle(context *Context) {
	fmt.Println("中件间2前置处理")
	context.RequestName = context.RequestName + "中间件2-"
	m.component.Handle(context)
	fmt.Println("中件间2后置处理")
}

type Middleware3 struct {
	BaseMiddleware
}

func (m *Middleware3) Handle(context *Context) {
	fmt.Println("中件间3前置处理")
	context.RequestName = context.RequestName + "中间件3-"
	m.component.Handle(context)
	fmt.Println("中件间3后置处理")
}

//go的继承，父类接受子类参数，只能通过实现同接口
//type parent interface {
//	test()
//}
//
//type A struct {
//}
//
//func (a *A) test() {
//	fmt.Println("我是AAAAA")
//}
//
//type B struct {
//	A
//}
//
////func (b *B) test() {
////	fmt.Println("我是BBBB")
////}
//
//func TestAb(p parent) {
//	p.test()
//}
