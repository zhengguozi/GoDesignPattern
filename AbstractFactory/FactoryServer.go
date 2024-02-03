package AbstractFactory

import "fmt"

type AbstractFactory interface {
	MakeComputer() FactoryComputer
	MakePhone() FactoryPhone
}

type XiaomiFactory struct {
}

func (base XiaomiFactory) MakeComputer() FactoryComputer {
	return &XiaomiComputer{
		BaseComputer{
			name: "小米电脑",
			size: 16,
		},
	}
}

func (base XiaomiFactory) MakePhone() FactoryPhone {
	return &XiaomiPhone{
		BasePhone{
			name: "小米手机14",
			size: 10,
		},
	}
}

type HuaweiFactory struct {
}

func (base HuaweiFactory) MakeComputer() FactoryComputer {
	return &HuaweiComputer{
		BaseComputer{
			name: "华为电脑",
			size: 17,
		},
	}
}

func (base HuaweiFactory) MakePhone() FactoryPhone {
	return &HuaweiPhone{
		BasePhone{
			name: "华为手机mate60Pro",
			size: 11,
		},
	}
}

func FactoryHelper(name string) AbstractFactory {
	if name == "mi" {
		return XiaomiFactory{}
	} else if name == "huawei" {
		return HuaweiFactory{}
	}
	return nil
}

func UseFactory(f AbstractFactory) {

	fmt.Println("不知名工厂生产的不知名电脑")
	absComputer1 := f.MakeComputer()
	fmt.Println(absComputer1.GetName())
	fmt.Println(absComputer1.GetSize())

	fmt.Println("不知名工厂生产的不知名手机")
	absPhone1 := f.MakePhone()
	fmt.Println(absPhone1.GetName())
	fmt.Println(absPhone1.GetSize())

}
