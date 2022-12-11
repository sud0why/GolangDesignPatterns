package main

import "fmt"

//需要目标接口的客户类不能直接重用适配类，因为它的接口不符合目标接口。
//客户端通过适配器类工作，该适配器类根据适配器实现目标接口
//适配器方式通过在运行时委托给适配对象（adaptee.specificOperation()）来实现目标接口。

//事后补救策略
//如果无必要，不要使用适配器

type OldLogin interface {
	LoginWithSocket()
}

type OldLoginObj struct {
}

func (l OldLoginObj) LoginWithSocket() {
	fmt.Println("LoginWithSocket")
}

// 能够login with file

type NewLoginObj struct {
}

func (l NewLoginObj) LoginWithFile() {
	fmt.Println("LoginWithFile")
}

// 适配器

type LoginAdapter struct {
	loginObj NewLoginObj
}

func (la LoginAdapter) LoginWithSocket() {
	la.loginObj.LoginWithFile()
}

func NewLoginAdapter() OldLogin {
	return &LoginAdapter{loginObj: NewLoginObj{}}
}

//业务

func main() {
	oldLoginObj := new(OldLoginObj)
	oldLoginObj.LoginWithSocket()

	newLoginObj := NewLoginAdapter()
	newLoginObj.LoginWithSocket()
}
