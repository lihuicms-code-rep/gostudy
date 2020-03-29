package main

import "fmt"

//抽闲层,实现层,业务逻辑层
//依赖倒转原则
//抽象层-->实现层-->业务逻辑层
//实现层向上只考虑抽象层提供了哪些
//业务逻辑层向下只考虑抽象层提供了哪些



//抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//实现层
type BenzCar struct {

}

func (bc *BenzCar) Run() {
	fmt.Println("奔驰车跑起来了......")
}

type BMWCar struct {

}

func (bc *BMWCar) Run() {
	fmt.Println("宝马车跑起来了......")
}


type ZHANG struct {

}

func (z3 *ZHANG) Drive(car Car) {
	fmt.Println("张三开车")
	car.Run()
}


type LI struct {

}

func (l4 *LI) Drive(car Car) {
	fmt.Println("李四开车")
	car.Run()
}



//业务逻辑
func main() {
	//1.张三开宝马
	var bcar Car
	bcar = &BMWCar{}

	var z3 ZHANG
	z3.Drive(bcar)
}

