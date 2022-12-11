package main

import "fmt"

//原型模式
//类实现Clone方法，克隆自身

type Data struct {
	Name string
	ID   int
}

func (d *Data) Clone() *Data {
	return &Data{Name: d.Name, ID: d.ID}
}

func main() {
	data1 := &Data{Name: "qwe", ID: 1}
	fmt.Println(*data1)
	data2 := data1.Clone()
	fmt.Println(*data2)
	if data1 != data2 {
		fmt.Println("Clone Success")
	}
}
