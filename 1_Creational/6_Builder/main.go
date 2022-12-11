package main

//建造者模式
//创建对象步骤多且复杂

import "fmt"

// AbsBuilder 抽象的Builder，定义需要实现的方法
type AbsBuilder interface {
	SetName(string)
	SetId(int)
	SetTest(bool)
}

// Director 管理类，
type Director struct {
	builder AbsBuilder
}

// NewDirector Director 的构造函数
func NewDirector(builder *Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 建造方法
func (d *Director) Construct(name string, id int, test bool) {
	d.builder.SetName(name)
	d.builder.SetId(id)
	d.builder.SetTest(test)
}

// Builder 实现
type Builder struct {
	name string
	id   int
	test bool
}

func (b *Builder) SetName(name string) {
	fmt.Println("SetName")
	b.name = name
}

func (b *Builder) SetId(id int) {
	fmt.Println("SetId")
	b.id = id
}

func (b *Builder) SetTest(test bool) {
	fmt.Println("SetTest")
	b.test = test
}

//业务代码

func main() {
	builder := new(Builder)
	director := NewDirector(builder)
	director.Construct("MyName", 666, true)
	fmt.Println(*builder)
}
