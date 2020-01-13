package main

import (
	"errors"
	"fmt"
)

//函数Go语言中支持函数、匿名函数和闭包，并且函数在Go语言中属于“一等公民”

//函数定义
func basic81() {

}

//可变参数**
func intSum2(x... int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

//返回值命名,函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

//定义函数类型，type calculation func(int, int) int
//简单来说，凡是满足这个条件的函数都是calculation类型的函数
type calculation func(int, int) int
func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

//c是一个函数类型的变量
var c calculation = add     //add和sub都能赋值给calculation类型的变量



//高阶函数分为函数作为参数和函数作为返回值两部分

//calc1是一个高阶函数,它将一个函数作为参数
func calc1(x, y int, op func(int, int) int) int {
	return op(x, y)
}

//do是一个高阶函数,它将函数作为返回值
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}

//闭包,闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境
//函数adder以func作为返回值
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

//defer语句
//由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等
//defer执行时机
//在Go语言的函数中return语句在底层并不是原子操作，它分为给返回值赋值和RET指令两步。
// 而defer语句执行的时机就在返回值赋值操作后，RET指令执行前
func defer1() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end" )
	//打印结果:start,end,3,2,1(defer语句按定义的逆序进行执行)
}

//defer经典案列
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		fmt.Println("defer语句执行时机....")
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calcdefer(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


//panic/recover
func funcA() {
	fmt.Println("func A")
}

//recover()必须搭配defer使用。
//defer一定要在可能引发panic的语句之前定义。
func funcB() {
	defer func() {
		err := recover()
		//如果程序出出现了panic错误,可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
		fmt.Println("recover process")
	}()
	panic("panic in B")
}


func funcC() {
	fmt.Println("func C")
}




func main() {
	ret1 := calc1(10, 20, add) //函数作为参数
	fmt.Println(ret1) //30

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	//闭包
	var f = adder()      //f是一个函数
	fmt.Println(f(10))   //10
	fmt.Println(f(20))   //30

	var f22 = adder2(10)
	fmt.Println(f22(20))   //30
	fmt.Println(f22(10))   //40

	//defer
	defer1()

	fmt.Println(f1())      //5
	fmt.Println(f2())      //6
	fmt.Println(f3())      //5
	fmt.Println(f4())      //5

	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x := 1
	y := 2
	defer calcdefer("AA", x, calcdefer("A", x, y))
	x = 10
	defer calcdefer("BB", x, calcdefer("B", x, y))
	y = 20

	funcB()
}


