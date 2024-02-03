package AbstractFactory

type FactoryComputer interface {
	SetName(name string)
	SetSize(size int)
	GetName() string
	GetSize() int
}

type BaseComputer struct {
	name string
	size int
}

func (base *BaseComputer) SetName(name string) {
	base.name = name
}

func (base *BaseComputer) GetName() string {
	return base.name
}

func (base *BaseComputer) SetSize(size int) {
	base.size = size
}

func (base *BaseComputer) GetSize() int {
	return base.size
}

type XiaomiComputer struct {
	BaseComputer
}

type HuaweiComputer struct {
	BaseComputer
}

type FactoryPhone interface {
	SetName(name string)
	SetSize(size int)
	GetName() string
	GetSize() int
}

type BasePhone struct {
	name string
	size int
}

func (base *BasePhone) SetName(name string) {
	base.name = name
}

func (base *BasePhone) GetName() string {
	return base.name
}

func (base *BasePhone) SetSize(size int) {
	base.size = size
}

func (base *BasePhone) GetSize() int {
	return base.size
}

type XiaomiPhone struct {
	BasePhone
}

type HuaweiPhone struct {
	BasePhone
}
