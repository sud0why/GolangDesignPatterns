package main

import "fmt"

//单例模式
//保证一个类永远只有一个对象，并且该对象的功能依然能被其他模块使用
//饿汉写法，永远占用内存

// 类必须是非公有的，不导出，不会被外部调用
type singleton struct {
}

// 指针指向唯一对象，不能改变方向，也不导出
var instance *singleton = new(singleton)

// GetInstance 提供导出函数用于获取singleton
// todo: 导出函数返回一个不导出的对象指针？
func GetInstance() *singleton {
	return instance
}

func (s *singleton) Do() {
	fmt.Println("singleton Do")
}

func main() {

	s1 := GetInstance()
	s2 := GetInstance()
	s1.Do()
	s2.Do()
	if s1 == s2 {
		fmt.Println("Same singleton")
	}
}
