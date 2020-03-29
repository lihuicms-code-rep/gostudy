package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}

//重写Phone的方法
func (nokia *NokiaPhone) call() {
	fmt.Println("i am nokia phone")
}


type ApplePhone struct {

}

func (apple *ApplePhone) call() {
	fmt.Println("i am apple phone")
}


//框架层,本身不知道哪个类型的phone来调用
//所以这里可以体现多态
//框架层不要改变,具体实现层可以改变
func callPhone(phone Phone) {
	phone.call()
}


func main() {
	//多态
	var phone Phone
	phone = &NokiaPhone{}
	phone.call()

	phone = &ApplePhone{}
	phone.call()

	callPhone(phone)

}
