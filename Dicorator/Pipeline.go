package Dicorator

type Pipeline struct {
	passable        *Context //这个应该理解为每个handle的函数参数才对，上下文，context
	middlewares     []Middleware
	middlewareChain Middleware
}

func (p *Pipeline) Send(passable *Context) *Pipeline {
	p.passable = passable
	return p
}

func (p *Pipeline) Through(middlewares ...Middleware) *Pipeline {
	p.middlewares = middlewares
	return p
}

// 传入业务回调组件,启动
func (p *Pipeline) Then(businessCallback *ConcreteBusinessComponent) {

	if len(p.middlewares) > 0 {
		p.middlewares[len(p.middlewares)-1].setComponent(businessCallback)
		for i := len(p.middlewares) - 2; i >= 0; i-- {
			p.middlewares[i].setComponent(p.middlewares[i+1])
		}
		p.middlewareChain = p.middlewares[0]
		p.middlewareChain.Handle(p.passable)
	}

}
