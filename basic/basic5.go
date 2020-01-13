package main

import "fmt"

//数组

//定义与初始化
func basic51() {
	//数组定义
	var a [3]int
	var b [4]int
	fmt.Println(a, b)

	//数组初始化
	var c [2]bool
	var d = [3]int{1,2,3}                           //1.使用初始化列表
	var cityArr = [...]string{"peking", "shenzhen"} //2.不写长度,编译器去推断个数
	var f = [...]string{1:"london", 3:"wuhan"}      //3.指定索引的值
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(cityArr)
	fmt.Println(f)
}

//遍历
func basic52() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

//多维数组
func basic53() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

//数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}



func main() {
	//basic51()
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}