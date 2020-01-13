package main

import "fmt"
//变量和常量

//变量声明
func basic1() {
	//1.标准声明
	var name string
	var age int
	//2.批量声明
	var (
		a string
		b int
		c bool
		d float64
	)

	fmt.Println(name, age, a, b, c, d)
}


//变量初始化
func basic2() {
	var name string = "lihui"   //提示type can be omitted
	var age int = 15
	var a, b = "lihui", 15   //一次初始化多个变量,同时类型推导
	fmt.Println(name, age, a, b)
}

//短变量声明
var m = 100     //全局变量
func basic3() {
	n := 100
	m := "llll" //局部变量
	fmt.Println(m, n)
}


//匿名变量
func foo() (int, string) {
	return 10, "lihui"
}

func basic4() {
	x, _ := foo()
	_, y := foo()
	fmt.Println(x, y)
	fmt.Println(pi, e, name, age, apple)
}

//常量
const pi = 3.14
const e = 2.71
const (
	name = "string"
	age          //声明多个变量时,如果省略了,表示和上面一行的值相同
	apple
)


//iota:常量计数器,在const出现时被重置为0，iota可以理解为const语句块中的行索引
//定义枚举时使用,这里可以理解为从iota再开始写就从行索索引开始算
const (
	n1 = iota           //0
	n2                  //1
	n3                  //2
	n4 = iota + 1       //4,中间插队的情况
	n5                  //5
	n6 = iota + 1       //6
	n7 = iota + 2       //8
)

//常见的用法
//1.可跳过的值
const (
	OutPut1 = iota     //0
	OutPut2            //1
	OutPut3            //3
	_
	_
	OutPut4            //5
)

//2.位掩码
const (
	B1 = 1 << iota    //1 << 0,  00000001
	B2                //1 << 1,  00000010
	B3                //1 << 2,  00000100
	B4                //1 << 3,  00001000
)

//3.定义数量级
type ByteSize float64
const (
	_ = iota
	KB ByteSize = 1 << (10 * iota)    // 1 << (10 * 1)
	MB                       // 1 << (10 * 2)
	GB                       // 1 << (10 * 3)
	TB                       // 1 << (10 * 4)
	PB                       // 1 << (10 * 5)    //1125 8999 0684 2624
	EB                       // 1 << (10 * 6)    //这里就溢出int(int64,int在64位机器上就是int64)的最大范围了2^64-1(922 3372 0368 5477 5807)
	ZB                       // 1 << (10 * 7)
	YB                       // 1 << (10 * 8)
)

//4.定义在一行的情况,iota这里我简单理解成还是按照行下标
const (
	Apple, Banana = iota + 1, iota + 2   //1, 2
	Pig, Dog                             //2, 3
	Chicken, lili                        //3, 4
)

func basic5() {
	fmt.Println(n1, n2, n3, n4, n5, n6, n7)
	fmt.Println(OutPut1, OutPut2, OutPut3, OutPut4)
	fmt.Println(B1, B2, B3, B4)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
	fmt.Println(Apple, Banana, Pig, Dog, Chicken, lili)
}

func main() {
	basic1()
	basic2()
	basic3()
	basic4()
	basic5()
}
