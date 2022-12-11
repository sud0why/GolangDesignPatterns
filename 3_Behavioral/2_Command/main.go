package main

import "fmt"

//命令模式
//比较容易设计成一个命令队列

// Command 菜单接口
type Command interface {
	Make()
}

// CommandA 菜单对应的厨师
type CommandA struct {
	receiver Receiver
}

func (c CommandA) Make() {
	fmt.Println("CommandA Make")
	c.receiver.MakeA()
}

// CommandB 菜单对应的厨师
type CommandB struct {
	receiver Receiver
}

func (c CommandB) Make() {
	fmt.Println("CommandB Make")
	c.receiver.MakeB()
}

// Receiver 厨师
type Receiver interface {
	MakeA()
	MakeB()
}

type ReceiverA struct {
}

func (r ReceiverA) MakeA() {
	fmt.Println("MakeA")
}

func (r ReceiverA) MakeB() {
	fmt.Println("MakeB")
}

// Invoker 服务员

type Invoker struct {
	CmdList []Command
}

func (i Invoker) Notify() {
	if i.CmdList != nil {
		for _, cmd := range i.CmdList {
			cmd.Make()
		}
		return
	} else {
		return
	}
}

//业务

func main() {
	receiver := new(ReceiverA)
	cmdA := CommandA{
		receiver: receiver,
	}
	cmdB := CommandB{
		receiver: receiver,
	}
	invoker := new(Invoker)
	invoker.CmdList = append(invoker.CmdList, cmdA)
	invoker.CmdList = append(invoker.CmdList, cmdB)
	invoker.Notify()
}
