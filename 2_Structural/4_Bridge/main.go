package main

import "fmt"

//桥接模式
//把抽象和实现分开
//把有关系的两个类分开独立存在，继承=>关联

//两个英雄，两个装备

type Hero interface {
	UseTool()
}

type Tool interface {
	Use()
}

//实现

type HeroA struct {
	tool Tool
}

func (h *HeroA) UseTool() {
	fmt.Println("HeroA")
	h.tool.Use()
}

func (h *HeroA) SetTool(tool Tool) {
	h.tool = tool
}

func NewHeroA() *HeroA {
	return &HeroA{}
}

type HeroB struct {
	tool Tool
}

func (h *HeroB) UseTool() {
	fmt.Println("HeroB")
	h.tool.Use()
}

func (h *HeroB) SetTool(tool Tool) {
	h.tool = tool
}

func NewHeroB() *HeroB {
	return &HeroB{}
}

// Tool 实现

type ToolA struct {
}

func (t ToolA) Use() {
	fmt.Println("Use ToolA")
}

type ToolB struct {
}

func (t ToolB) Use() {
	fmt.Println("Use ToolB")
}

//业务

func main() {
	heroA := NewHeroA()
	toolA := new(ToolA)
	heroA.SetTool(toolA)
	heroA.UseTool()

	heroB := NewHeroB()
	heroB.SetTool(toolA)
	heroB.UseTool()
	toolB := new(ToolB)
	heroB.SetTool(toolB)
	heroB.UseTool()

}
