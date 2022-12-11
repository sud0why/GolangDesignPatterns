package main

import "fmt"

//享元模式
//共用对象,节省内存资源
//工厂中存放一个map
//连接池,对象池的实现

//享元工厂

type UnitFactory struct {
	maps map[string]*Unit
}

func NewUnitFactory() *UnitFactory {
	return &UnitFactory{
		maps: make(map[string]*Unit),
	}
}

func (uf *UnitFactory) Get(name string) *Unit {
	unit := uf.maps[name]
	if unit == nil {
		unit = &Unit{
			name: name,
			data: "default",
		}
		uf.maps[name] = unit
		return unit
	}
	return unit
}

// 享元

type Unit struct {
	name string
	data string
}

func (u *Unit) Show() {
	fmt.Println(u.name + "\t" + u.data)
}

func (u *Unit) SetData(data string) {
	u.data = data
}

var UF = NewUnitFactory()

// 业务

func init() {
	UF.maps["qqq"] = &Unit{
		name: "qqq",
		data: "init",
	}
	UF.maps["www"] = &Unit{
		name: "www",
		data: "init",
	}
}

func main() {
	qUnit := UF.Get("qqq")
	qUnit.Show()
	qUnit.SetData("change")
	qUnit.Show()

	eUnit := UF.Get("eee")
	eUnit.Show()

	qUnit2 := UF.Get("qqq")
	fmt.Println(qUnit2 == qUnit)
}
