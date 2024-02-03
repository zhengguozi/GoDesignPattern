package Builder

//生成器模式，同个产品构造复杂，最终表现形式有点差异

type Car struct {
	Brand                 string
	Color                 string
	EngineType            string
	EnableSeatVentilation bool //座椅通风
	EnableSeatHeating     bool //座椅加热
}

type Option func(car *Car)

// 生成器接口
type Builder interface {
	SetBrand(v string) Option
	SetColor(v string) Option
	SetEngineType(v string) Option
	SetEnableSeatVentilation(v bool) Option
	SetEnableSeatHeating(v bool) Option

	NewCar(options ...Option) *Car
}

// 基础builder,限制了通过自定义Option来生成car, 具体的builder只需要实现各自的Option即可
type BaseBuilder struct {
}

func (base BaseBuilder) NewCar(options ...Option) *Car {
	myCar := &Car{}
	for _, optionItem := range options {
		optionItem(myCar)
	}
	return myCar
}

// 小米汽车builder
type XiaomiBuilder struct {
	BaseBuilder
}

//func (builder XiaomiBuilder) SetBrand() Option {
//	return func(car *Car) {
//		cat.s
//	}
//}
