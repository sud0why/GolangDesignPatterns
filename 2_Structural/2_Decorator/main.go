package main

import "fmt"

//装饰模式
//装饰模式就是将对象封装到另一个对象中，为原对象绑定新的行为。
//类似代理模式，但是代理模式主要是添加原始类无关的功能，装饰器用于添加相关的功能。

// Phone 抽象手机
type Phone interface {
	Show()
}

// PhoneDecorator 手机装饰器
type PhoneDecorator struct {
	phone Phone
}

func (d *PhoneDecorator) Show() {
	d.phone.Show()
}

// Phone的实现

type HuaweiPhone struct {
}

func (hw *HuaweiPhone) Show() {
	fmt.Println("展示华为手机")
}

type XiaomiPhone struct {
}

func (xm *XiaomiPhone) Show() {
	fmt.Println("展示华为手机")
}

// 具体装饰器

type MoDecorator struct {
	PhoneDecorator
}

func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("加膜")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{PhoneDecorator{
		phone: phone,
	}}
}

type KeDecorator struct {
	PhoneDecorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("加壳")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{PhoneDecorator{
		phone: phone,
	}}
}

//业务

func main() {
	huawei := new(HuaweiPhone)
	huawei.Show()

	keHuawei := NewKeDecorator(huawei)
	keHuawei.Show()

	kekeHuawei := NewKeDecorator(keHuawei)
	kekeHuawei.Show()

	mokekeHuawei := NewMoDecorator(kekeHuawei)
	mokekeHuawei.Show()
}
