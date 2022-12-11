package main

import (
	"fmt"
	"sync"
)

//单例模式
//保证一个类永远只有一个对象，并且该对象的功能依然能被其他模块使用
//lazy写法，用时初始化

// 类必须是非公有的，不导出，不会被外部调用
type singleton struct {
}

var once sync.Once

// 指针指向唯一对象，不能改变方向，也不导出
var instance *singleton = new(singleton)

// GetInstance 提供导出函数用于获取singleton
func GetInstance() *singleton {
	//底层实现也是有一个标志位，原子操作来判断
	//或者加锁
	once.Do(func() {
		instance = new(singleton)
	})
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
