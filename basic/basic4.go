package main

import "fmt"

//流程控制

//if else分支结构
func basic41() {
	score := 65
	if score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	//特殊写法
	if score1 := 65; score1 >= 90 {
		fmt.Println("A")
	} else if score1 > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}


//for循环结构
func basic42() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//for range 键值循环
	//数组、切片、字符串返回索引和值
	//map返回键和值
	//通道（channel）只返回通道内的值
}

//switch case
func basic43() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}

	//一个分支可以有多个值
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//分支还可以使用表达式
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}

//goto 调到指定标签
// goto语句通过标签进行代码间的无条件跳转。
// goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。
// Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时
func gotoDemo1() {
	var breakFlag bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				breakFlag = true
				break
			}
			fmt.Printf("%v-%v\n", i, j)
		}
		// 外层for循环判断
		if breakFlag {
			break
		}
	}
}

func gotoDemo2() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
}

func main() {
	gotoDemo1()
	gotoDemo2()
}