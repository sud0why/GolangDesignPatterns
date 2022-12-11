package main

import "fmt"

//中介者模式
//聊天室，很多人，一个聊天室
//一个人发消息，所有的人都要收到通知
//中介者模式将聊天室作为中介，人只需要和聊天室耦合，发送完消息，调用中介（聊天室）的方法通知别人

type User struct {
	Name string
	Room *Room
}

func (u User) SendMsg(data string) {
	u.Room.Notify(u, data)
}

func (u User) RecvMsg(data string) {
	fmt.Printf("User %s recv msg: %s\n", u.Name, data)
}

type Room struct {
	Users []User
}

func (r Room) Notify(u User, data string) {
	for _, user := range r.Users {
		if user.Name != u.Name {
			user.RecvMsg(fmt.Sprintf("%s say %s", u.Name, data))
		}
	}
}

func main() {
	chatroom := Room{}
	u1 := User{
		Name: "zhangsan",
		Room: &chatroom,
	}
	u2 := User{
		Name: "lisi",
		Room: &chatroom,
	}
	u3 := User{
		Name: "wangwu",
		Room: &chatroom,
	}
	chatroom.Users = append(chatroom.Users, []User{u1, u2, u3}...)
	u1.SendMsg("Hello")
	u2.SendMsg("Bye")
}
