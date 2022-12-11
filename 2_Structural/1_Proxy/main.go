package main

import "fmt"

//代理模式
//抽象接口
//代理也实现了抽象接口，在使用或者访问对象之前进行检查等访问控制
//代理模式通过代理来替代真实服务，通常代理类与真实类具有相同的接口，在代理类中可以做一些额外操作（访问控制、缓存等）
//比如web框架，用户访问接口时进行检查和日志记录等

//动态代理参考：https://lailin.xyz/post/proxy.html

// Page 抽象层
type Page interface {
	Watch() string
}

//实现

type APage struct {
}

func (a APage) Watch() string {
	fmt.Println("APage Watch")
	return "APage Content"
}

type BPage struct {
}

func (b BPage) Watch() string {
	fmt.Println("BPage Watch")
	return "BPage Content"
}

//代理实现

type PageProxy struct {
	page     Page
	userInfo UserInfo
}

type UserInfo struct {
	Role string
}

func (p PageProxy) Check() bool {
	if p.userInfo.Role == "Admin" {
		fmt.Println("Admin can watch")
		return true
	} else {
		fmt.Println("Non Admin can't watch")
		return false
	}
}

func (p PageProxy) Log(content string) {
	fmt.Println(content)
}

func (p PageProxy) Watch() string {
	if p.Check() {
		content := p.page.Watch()
		p.Log(content)
		return content
	} else {
		return ""
	}
}

func NewPageProxy(page Page, userinfo UserInfo) *PageProxy {
	return &PageProxy{page: page, userInfo: userinfo}
}

//业务实现

func main() {
	//管理员用户
	user1 := UserInfo{
		Role: "Admin",
	}
	apage := APage{}
	proxy1 := NewPageProxy(apage, user1)
	proxy1.Watch()

	//普通用户
	user2 := UserInfo{
		Role: "User",
	}
	bpage := BPage{}
	proxy2 := NewPageProxy(bpage, user2)
	proxy2.Watch()
}
