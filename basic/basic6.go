package main

import "fmt"

//切片
//切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
//切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合

func basic61() {
	//切片定义
	var a []string
	var b = []bool{true, false}    //定义并初始化
	fmt.Println(a == nil)
	fmt.Println(b == nil)

	//基于数组定义切片
	c := [5]int{1, 2, 3, 4, 5}
	d := c[1:4]
	fmt.Printf("type of d is %T\n", d)

	//切片再切片
	//这里len求出的就是实际的元素个数,而cap的求解: 原数组或者源切片的cap - start
	a1 := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a1:%v type:%T len:%d  cap:%d\n", a1, a1, len(a1), cap(a1))
	b1 := a1[1:3]
	fmt.Printf("b1:%v type:%T len:%d  cap:%d\n", b1, b1, len(b1), cap(b1))  //cap(b1) = cap(a1)-1 = 6-1= 5
	c1 := b1[1:5]
	fmt.Printf("c1:%v type:%T len:%d  cap:%d\n", c1, c1, len(c1), cap(c1))  //cap(c1) = cap(b1)-1 = 5-1=4

	//使用make构造切片
	f := make([]int, 2, 10)
	fmt.Printf("f:%v type:%T len:%d  cap:%d\n", f, f, len(f), cap(f))

	//切片不能直接比较,唯一合法的比较操作是和nil比较，所以要判断一个切片是否是空的，要是用len(s) == 0来判断，不应该使用s == nil来判断


	//切片的赋值拷贝
	s1 := make([]int, 3) //[0 0 0]
	s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
	s2[0] = 100
	fmt.Println(s1) //[100 0 0]
	fmt.Println(s2) //[100 0 0]

	//切片遍历

	//append()方法为切片添加元素
	//Go语言的内建函数append()可以为切片动态添加元素，每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
	// 当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
	// “扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append函数的返回值

	var numSlice []int
	for i := 0; i < 10; i++ {
		numSlice = append(numSlice, i)
		fmt.Printf("%v  len:%d  cap:%d  ptr:%p\n", numSlice, len(numSlice), cap(numSlice), numSlice)
	}

	//append()函数还支持一次性追加多个元素
	var citySlice []string
	// 追加一个元素
	citySlice = append(citySlice, "北京")
	// 追加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")
	// 追加切片
	aa := []string{"成都", "重庆"}
	citySlice = append(citySlice, aa...)
	fmt.Println(citySlice) //[北京 上海 广州 深圳 成都 重庆]


	//切片的扩容策略
	//查看源码:runtime/slice.go中growslice方法的实现


	//使用copy()函数复制切片
	a11 := []int{1, 2, 3, 4, 5}
	c11 := make([]int, 5, 5)
	copy(c11, a11)     //使用copy()函数将切片a11中的元素复制到切片c11
	fmt.Println(a11) //[1 2 3 4 5]
	fmt.Println(c11) //[1 2 3 4 5]
	c11[0] = 1000
	fmt.Println(a11) //[1 2 3 4 5]
	fmt.Println(c11) //[1000 2 3 4 5]

	//从切片中删除元素,Go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	a22 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为2的元素
	a22 = append(a22[:2], a22[3:]...)
	fmt.Println(a22) //[30 31 33 34 35 36 37]
}



func main() {
	basic61()
}