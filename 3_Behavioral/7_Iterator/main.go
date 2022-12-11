package main

import "fmt"

//迭代器模式

type Iterator interface {
	HasMore() bool
	Next() Member
}

type Iterable interface {
	GetIterator() Iterator
}

type Student struct {
	Name string // 姓名
}

// Member 空接口，用于统一迭代器Next的返回值
type Member interface {
}

// NewStudent 创建学生对象
func NewStudent(name string) *Student {
	return &Student{
		Name: name,
	}
}

func (t *Student) String() string {
	return t.Name
}

// ClassIterator 教室容器的迭代器
type ClassIterator struct {
	Class *Class
	Index int
}

func (c *ClassIterator) Next() Member {
	if !c.HasMore() {
		panic("Iterator Out Of Index")
	}
	stu := c.Class.Stus[c.Index]
	c.Index++
	return stu
}

func (c *ClassIterator) HasMore() bool {
	return c.Index < len(c.Class.Stus)
}

// Class 教室容器类
type Class struct {
	Stus []Student
}

func (c *Class) GetIterator() Iterator {
	return &ClassIterator{Class: c, Index: 0}
}

func main() {
	class := Class{
		Stus: []Student{
			{Name: "QQQ"},
			{Name: "WWW"},
			{Name: "EEE"},
		},
	}
	classIterator := class.GetIterator()
	for classIterator.HasMore() {
		fmt.Println(classIterator.Next().(Student).Name)
	}
}
