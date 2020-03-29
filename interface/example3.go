package main

import "fmt"

//以前的情况描述

//奔驰
type Benz struct {

}

func (b *Benz) Run() {
	fmt.Println("benz is running")
}

//宝马
type BWM struct {

}

func (b *BWM) Run() {
	fmt.Println("bwm is running")
}


//张三司机
type Z3 struct {

}

func (z *Z3) DriveBenz(b *Benz) {
	fmt.Println("张三开奔驰......")
	b.Run()
}

func (z *Z3) DriveBWM(b *BWM) {
	fmt.Println("张三开宝马......")
	b.Run()
}

//李四司机
type L4 struct {

}

func (l *L4) DriveBenz(b *Benz) {
	fmt.Println("李四开奔驰......")
	b.Run()
}

func (l *L4) DriveBWM(b *BWM) {
	fmt.Println("李四开宝马......")
	b.Run()
}

func main() {
	//业务1:张三开奔驰
	benz := &Benz{}
	z3 := Z3{}
	z3.DriveBenz(benz)
}
