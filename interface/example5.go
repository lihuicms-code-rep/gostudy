package main

import (
	"fmt"
)

//依赖倒转原则的案列

//抽象层
//cpu
type CPU interface {
	Calculate()
}

//内存
type Memory interface {
	Storage()
}

//显卡
type Card interface {
	Display()
}

type Computer struct {
	cpu CPU
	mem Memory
	card Card
}

//伴随一个new
//参数传递的是interface接口,实际上传递具体的实现
func NewComputer(cpu CPU, mem Memory, card Card) *Computer {
	return &Computer{
		cpu,
		mem,
		card,
	}
}

func (c *Computer) Work() {
	c.cpu.Calculate()
	c.card.Display()
	c.mem.Storage()
}

//----实现层

type IntelCPU struct {

}

func (ic *IntelCPU) Calculate() {
	fmt.Println("intel cpu run")
}

type IntelCard struct {

}

func (ic *IntelCard) Display() {
	fmt.Println("intel card run")
}

type IntelMemory struct {

}

func (im *IntelMemory) Storage() {
	fmt.Println("intel memory run")
}

type KingstonMemeory struct {

}

func (km *KingstonMemeory) Storage() {
	fmt.Println("kingston memory run")
}

type NavidiaCard struct {

}

func (nc *NavidiaCard) Display() {
	fmt.Println("navidia card display")
}





//业务逻辑层
func main() {
	//1.组成一台Intel系列的电脑,用intel的cpu,内存,显卡
	compute1 := NewComputer(&IntelCPU{}, &IntelMemory{}, &IntelCard{})
	compute1.Work()

	//2.组成一台intel的CPU,Kingston的内存,NVIDIA的显卡
	computer2 := NewComputer(&IntelCPU{}, &KingstonMemeory{}, &NavidiaCard{})
	computer2.Work()
}
