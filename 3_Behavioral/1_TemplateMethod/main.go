package main

import "fmt"

//模板方法模式
//两件事情的流程相似

type Beverage interface {
	step1()
	step2()
	step3()
	step4()
}

// go没办法在接口中定义实现流程，所以只能封装一个类

type Template struct {
	b Beverage
}

func (t *Template) Make() {
	if t == nil {
		return
	}
	t.b.step1()
	t.b.step2()
	t.b.step3()
	t.b.step4()
}

// 实现

type Tea struct {
	Template
}

func (t Tea) step1() {
	fmt.Println("Tea step1")
}

func (t Tea) step2() {
	fmt.Println("Tea step2")
}

func (t Tea) step3() {
	fmt.Println("Tea step3")
}

func (t Tea) step4() {
	fmt.Println("Tea step4")
}

type Coffee struct {
	Template
}

func (c Coffee) step1() {
	fmt.Println("Coffee step1")
}

func (c Coffee) step2() {
	fmt.Println("Coffee step2")
}
func (c Coffee) step3() {
	fmt.Println("Coffee step3")
}
func (c Coffee) step4() {
	fmt.Println("Coffee step4")
}

func main() {
	tea := new(Tea)
	//给模板赋值
	tea.b = tea
	tea.Make()

	coffee := new(Coffee)
	//给模板赋值
	coffee.b = coffee
	coffee.Make()
}
