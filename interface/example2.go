package main

import "fmt"

//开闭原则的体现

//之前的设计,如果一味在一个类上添加新功能
//类会越来越臃肿,耦合度会很高

//改造
//面向对象的开闭原则
//通过额外添加功能而不是修改原来的类

//抽象的业务员Banker
type AbstarctBanker interface {
	DoBusiness()        //业务接口
}

//底下取划分成各个不同功能的Banker
//1.专门实现存款业务的Banker
type SaveBanker struct {

}

func (sb *SaveBanker) DoBusiness() {
	fmt.Println("save banker")
}



//2.专门实现转账业务的Banker
type TransBanker struct {

}

func (tb *TransBanker) DoBusiness() {
	fmt.Println("trans banker")
}

//此时,额外添加另外的功能
type PayBanker struct {

}

func (pb *PayBanker) DoBusiness() {
	fmt.Println("pay banker")
}

//每一个具体的功能或者说系统都可以单独的来写


//框架层提供一个抽象的接口
//这样不用考虑具体是谁
func BankerBusiness(banker AbstarctBanker) {
	banker.DoBusiness()
}



func main() {
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransBanker{})
	BankerBusiness(&PayBanker{})
}

