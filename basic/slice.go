package main

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
	"unsafe"
)

//切片
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}

func main() {
	x := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("x type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x, &x, len(x), cap(x), x)

	//基于数组的切片
	x1 := x[:]
	fmt.Printf("x1 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x1, x1, len(x1), cap(x1), x1)

	x2 := x[2:5]
	fmt.Printf("x2 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x2, x2, len(x2), cap(x2), x2)

	x3 := x[2:5:7]
	fmt.Printf("x3 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x3, x3, len(x3), cap(x3), x3)

	x4 := x[4:]
	fmt.Printf("x4 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x4, x4, len(x4), cap(x4), x4)

	x5 := x[:4]
	fmt.Printf("x5 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x5, x5, len(x5), cap(x5), x5)

	x6 := x[:4:6]
	fmt.Printf("x6 type:%T\t addr:%p\t len:%d\t cap:%d\t value:%+v\n", x6, x6, len(x6), cap(x6), x6)

	//可以直接创建切片对象,无需预先准备数组,借助make函数或者显示的初始化语句
	//它会自动完成底层数组的内存分配
	s1 := make([]int, 3, 5)
	s2 := make([]int, 3)
	s3 := []int{10, 20, 5: 30}

	fmt.Printf("s1:%+v, len:%d, cap:%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2:%+v, len:%d, cap:%d\n", s2, len(s2), cap(s2))
	fmt.Printf("s3:%+v, len:%d, cap:%d\n", s3, len(s3), cap(s3))

	//注意
	var a []int      //仅仅定义了一个切片类型a,未初始化
	b := []int{}

	fmt.Println(a == nil, b == nil)  //true, false

	//输出详细的信息
	fmt.Printf("a:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)))
	fmt.Printf("b:%#v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)))

	//切片不支持==比较运算
	c := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("c value:%+v, addr:%p\n", c, c)
	fmt.Printf("c0 value:%+v, addr:%p\n", c[0], &c[0])
	p := &c
	p0 := &c[0]
	p1 := &c[1]

    fmt.Println(p, p0, p1)

	(*p)[0] += 100
	*p1 += 100

	fmt.Printf("c value:%+v, addr:%p\n", c, c)

	//如果元素的类型是切片,可以实现交替数组
	d := [][]int{
		{1, 2, 3},
		{4, 5},
		{6},
		{7, 8, 8, 10},
	}

	fmt.Printf("d type:%T\t addr:%p\t value:%+v\n", d, d, d)

	//再认识一次几个的区别
	var a1 [10]int
	b1 := make([]int, 10)
	c1 := make([]int, 0, 10)

	fmt.Printf("a1:%+v\t, b1:%+v\t c1:%+v\n", a1, b1, c1)

	//reslice,切片再切片
	f := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	f1 := f[3:7]
	fmt.Printf("f1:%+v\t addr:%p\t len:%d\t cap:%d\n", f1, f1, len(f1), cap(f1))
	f2 := f1[1:3]
	fmt.Printf("f2:%+v\t addr:%p\t len:%d\t cap:%d\n", f2, f2, len(f2), cap(f2))

	//用slice模拟栈的行为
	//stackfun(5)

	//append
	g1 := make([]int, 0, 100)
	g2 := g1[:2:4]
	g3 := append(g2, 1, 2, 3, 4, 5, 6)   //超过切片cap后,会按照cap*2进行扩容，重新分配空间

	fmt.Printf("g1:%+v, addr:%p, len:%d, cap:%d\n", g1, g1, len(g1), cap(g1))

	fmt.Printf("g2:%+v, addr:%p, len:%d, cap:%d\n", g2, g2, len(g2), cap(g2))

	fmt.Printf("g3:%+v, addr:%p, len:%d, cap:%d\n", g3, g3, len(g3), cap(g3))


	//还有这种写法下
	var h []int         //未初始化
	fmt.Println(h == nil)
	fmt.Printf("h:%+v, addr:%p, len:%d, cap:%d\n", h, h, len(h), cap(h))  //这里addr就是0x0
	h = append(h, 1, 2, 3, 4, 5)
	fmt.Printf("h:%+v, addr:%p, len:%d, cap:%d\n", h, h, len(h), cap(h))  //len:5, cap:6, 加一个扩容2,3个时,扩为6

	//copy

	i := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	i1 := i[5:8]
	n := copy(i[4:], i1)        //i1 copy 到 i[4:]
	fmt.Println(n, i)

	i2 := make([]int, 6)
	n = copy(i2, i)           //i copy到 i2,注意i2只有6个cap
	fmt.Println(n, i2)

	//直接将字符串复制到[]byte
	j := make([]byte, 3)
	n = copy(j, "李虎")
	fmt.Println(n, j)






}

//利用slice的reslice可以实现一个栈
func stackfun(c int) {
	stack := make([]int, 0, c)

	//入栈操作
	push := func(x int) error {
		n := len(stack)
		if n == c {
			return errors.New("stack is full")
		}

		stack = stack[:n+1]
		stack[n] = x

		return nil
	}

	//出栈操作
	pop := func() int {
		n := len(stack)
		if n == 0 {
			return -1
		}

		x := stack[n-1]
		stack = stack[:n-1]
		return x
	}

	for i := 0; i < 7; i++ {
		fmt.Printf("push %d:%v, cur stack:%+v, len:%d, cap:%d\n", i, push(i), stack, len(stack), cap(stack))
	}

	for i := 0; i < 7; i++ {
		fmt.Printf("pop %d:%v, cur stack:%+v, len:%d, cap:%d\n", i, pop(), stack, len(stack), cap(stack))
	}
}
