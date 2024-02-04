package Builder

import "fmt"

// 主管类,用来组织builder，这个其实是可选的，客户端也可以自己构建复杂的builder
type Director struct {
	MyBuilder Builder
}

func (d Director) SetBuilder(builder Builder) {
	d.MyBuilder = builder
}

func (d Director) MakeCar() *Car {
	return d.MyBuilder.NewCar(
		d.MyBuilder.SetColor("白"),
		d.MyBuilder.SetBrand("su7"),
		d.MyBuilder.SetEngineType("50"),
		d.MyBuilder.SetEnableSeatVentilation(false),
		d.MyBuilder.SetEnableSeatHeating(false),
	)
}

func (d Director) ShowCar(car *Car) {
	fmt.Println(*car)
}

func NewDiretor() Director {
	return Director{
		MyBuilder: XiaomiBuilder{
			BaseBuilder{},
		},
	}
}
