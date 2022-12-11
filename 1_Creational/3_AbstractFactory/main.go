package main

import "fmt"

// 抽象工厂
// 把产品也抽象了，并且有很多个工厂
// 抽象出的工厂的功能的特定的，比如生产某些产品
// 添加工厂，可以，符合开闭原则
// 添加产品，不符合开闭原则

// 抽象层

type AbsCard interface {
	Display()
}

type AbsCPU interface {
	Calculate()
}

type AbsMemory interface {
	Storage()
}

type AbsFactory interface {
	NewCard() AbsCard
	NewCPU() AbsCPU
	NewMemory() AbsMemory
}

// 实现层

type IntelMemory struct{}

func (i IntelMemory) Storage() {
	fmt.Println("IntelMemory Storage")
}

type IntelCard struct{}

func (i IntelCard) Display() {
	fmt.Println("IntelCard Display")
}

type IntelCPU struct{}

func (i IntelCPU) Calculate() {
	fmt.Println("IntelCPU Calculate")
}

type IntelFactory struct {
}

func (i IntelFactory) NewCard() AbsCard {
	return new(IntelCard)
}

func (i IntelFactory) NewCPU() AbsCPU {
	return new(IntelCPU)
}

func (i IntelFactory) NewMemory() AbsMemory {
	return new(IntelMemory)
}

type NvidiaMemory struct{}

func (i NvidiaMemory) Storage() {
	fmt.Println("NvidiaMemory Storage")
}

type NvidiaCard struct{}

func (i NvidiaCard) Display() {
	fmt.Println("NvidiaCard Display")
}

type NvidiaCPU struct{}

func (i NvidiaCPU) Calculate() {
	fmt.Println("NvidiaCPU Calculate")
}

type NvidiaFactory struct {
}

func (i NvidiaFactory) NewCard() AbsCard {
	return new(NvidiaCard)
}

func (i NvidiaFactory) NewCPU() AbsCPU {
	return new(NvidiaCPU)
}

func (i NvidiaFactory) NewMemory() AbsMemory {
	return new(NvidiaMemory)
}

type KingstonMemory struct{}

func (i KingstonMemory) Storage() {
	fmt.Println("KingstonMemory Storage")
}

type KingstonCard struct{}

func (i KingstonCard) Display() {
	fmt.Println("KingstonCard Display")
}

type KingstonCPU struct{}

func (i KingstonCPU) Calculate() {
	fmt.Println("KingstonCPU Calculate")
}

type KingstonFactory struct {
}

func (i KingstonFactory) NewCard() AbsCard {
	return new(KingstonCard)
}

func (i KingstonFactory) NewCPU() AbsCPU {
	return new(KingstonCPU)
}

func (i KingstonFactory) NewMemory() AbsMemory {
	return new(KingstonMemory)
}

// 业务代码

func main() {
	intelFactory := new(IntelFactory)
	intelCPU := intelFactory.NewCPU()
	intelCPU.Calculate()

	nvidiaFactory := new(NvidiaFactory)
	nvidiaCard := nvidiaFactory.NewCard()
	nvidiaCard.Display()

	kingstonFactory := new(KingstonFactory)
	kingstonMemory := kingstonFactory.NewMemory()
	kingstonMemory.Storage()
}
