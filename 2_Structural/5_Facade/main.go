package main

import "fmt"

//外观模式
//为子系统提供高层接口，提高易用性
//细粒度接口的包装

type API interface {
	Test()
}

type APIA struct{}

func NewAPIA() API {
	return &APIA{}
}

func (*APIA) Test() {
	fmt.Println("APIA Test")
}

type APIB struct{}

func NewAPIB() API {
	return &APIB{}
}

func (*APIB) Test() {
	fmt.Println("APIB Test")
}

// 外观

type APIALL struct {
	a API
	b API
}

func NewAPIALL() API {
	return &APIALL{
		a: NewAPIA(),
		b: NewAPIB(),
	}
}

func (a *APIALL) Test() {
	a.a.Test()
	a.b.Test()
}

func main() {
	apiALL := NewAPIALL()
	apiALL.Test()
}
