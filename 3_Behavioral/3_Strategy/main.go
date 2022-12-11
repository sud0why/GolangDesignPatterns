package main

import "fmt"

//策略模式
//常见
//运行时可替换算法，实现算法分离
//有一个文件服务，两种提供文件内容的策略，Json编码和Base64编码
//用户需要知道有哪些策略类可以使用

// 两种策略

type Strategy interface {
	GetData(string)
}

type JsonStrategy struct {
}

func (j JsonStrategy) GetData(data string) {
	fmt.Println("Json Data")
	fmt.Println(data)
}

type Base64Strategy struct {
}

func (b Base64Strategy) GetData(data string) {
	fmt.Println("Base64 Data")
	fmt.Println(data)
}

// 策略的应用

type Info struct {
	data     string
	strategy Strategy
}

func (i *Info) SetStrategy(strategy Strategy) {
	i.strategy = strategy
}

func (i *Info) GetData() {
	if i.strategy == nil {
		return
	}
	i.strategy.GetData(i.data)
}

//业务

func main() {
	info := Info{
		data: "qweasd",
	}
	info.SetStrategy(new(JsonStrategy))
	info.GetData()

	info.SetStrategy(new(Base64Strategy))
	info.GetData()
}
