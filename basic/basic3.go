package main

import "fmt"

//运算符

//算数运算符
func basic31() {
	fmt.Println(1 + 2)
	fmt.Println( 1 - 2)
	fmt.Println( 1.5 * 2)
	fmt.Println( 1 / 2.0)
	fmt.Println( 1 % 2)
}

//关系运算符
func basic32() {
	fmt.Println(0 == 0)
	fmt.Println(0 != 1)
	fmt.Println(1.2 > 2.4)
	fmt.Println(1.1 >= 1.1)
	fmt.Println(-0.3 < -0.2)
	fmt.Println(-1e4 <= -2e4)
}

//逻辑运算符
func basic33() {
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

//位运算符
func basic34() {
	fmt.Println(1&2)    //0001&0010=0000
	fmt.Println(1|2)    //0001|0010=0011
	fmt.Println(1^2)    //异或,0001^0010=0011
	fmt.Println(1<<2)   //左移两位,0001<<2=0100(简单理解为就是a*2^n)
	fmt.Println(1>>2)   //右移两位,0001>>2=0000(简单理解就是a/2^n)
}

//赋值运算符
func basic35() {
	var a uint32  = 1
	var b uint32 = 2
	a <<= b
	fmt.Println(a)
	//类似的还有:+=,-=,*=,/=,%=,<<=,>>=,&=,|=,^=
}

func main() {
	basic31()
	basic32()
	basic33()
	basic34()
	basic35()
}