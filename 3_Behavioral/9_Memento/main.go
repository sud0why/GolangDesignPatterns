package main

import (
	"fmt"
	"strings"
)

//备忘录模式
//原始对象Originator状态，Memento 用于储存状态的对象，

type Role struct {
	Name  string
	State string
}

type Memento struct {
	Data string
}

func (r *Role) Snapshot() Memento {
	return Memento{Data: fmt.Sprintf("%s,%s", r.Name, r.State)}
}

func (r *Role) Restore(m Memento) {
	data := strings.Split(m.Data, ",")
	r.Name = data[0]
	r.State = data[1]
}

func (r *Role) ChangeName(name string) {
	r.Name = name
}

func (r *Role) Die() {
	r.State = "Die"
}

func main() {
	role := Role{
		Name:  "old",
		State: "Live",
	}
	fmt.Println(role)
	snapshot := role.Snapshot()
	role.ChangeName("new")
	role.Die()
	fmt.Println(role)
	role.Restore(snapshot)
	fmt.Println(role)
}
