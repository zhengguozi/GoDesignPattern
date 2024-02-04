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

func (builder XiaomiBuilder) SetBrand(v string) Option {
	return func(car *Car) {
		car.Brand = "小米" + v
	}
}

func (builder XiaomiBuilder) SetColor(v string) Option {
	return func(car *Car) {
		car.Color = "小米" + v
	}
}

func (builder XiaomiBuilder) SetEngineType(v string) Option {
	return func(car *Car) {
		car.EngineType = "小米特制" + v
	}
}

func (builder XiaomiBuilder) SetEnableSeatVentilation(v bool) Option {
	return func(car *Car) {
		car.EnableSeatVentilation = v
	}
}

func (builder XiaomiBuilder) SetEnableSeatHeating(v bool) Option {
	return func(car *Car) {
		car.EnableSeatHeating = v
	}
}

// 华为问界汽车builder
type WenjieBuilder struct {
	BaseBuilder
}

func (builder WenjieBuilder) SetBrand(v string) Option {
	return func(car *Car) {
		car.Brand = "问界" + v
	}
}

func (builder WenjieBuilder) SetColor(v string) Option {
	return func(car *Car) {
		car.Color = "问界" + v
	}
}

func (builder WenjieBuilder) SetEngineType(v string) Option {
	return func(car *Car) {
		car.EngineType = "问界特制" + v
	}
}

func (builder WenjieBuilder) SetEnableSeatVentilation(v bool) Option {
	return func(car *Car) {
		car.EnableSeatVentilation = v
	}
}

func (builder WenjieBuilder) SetEnableSeatHeating(v bool) Option {
	return func(car *Car) {
		car.EnableSeatHeating = v
	}
}
