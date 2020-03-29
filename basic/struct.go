package main

import (
	"fmt"
	"os"
	"reflect"
)

//结构体将多个不同类型命名字段打包成一个符合类型
type node struct {
	_ int
	id int
	next *node     //可以使用自身指针成员类型
}

func main() {

	//命名初始化
	n1 := node {
		id : 1,
	}

	n2 := node {
		id :2,
		next : &n1,
	}

	fmt.Println(n1, n2)

	//定义匿名结构体变量
	u := struct {
		Name string
		Age int
	} {
		Name:"lihui",
		Age : 11,
	}

	fmt.Println(u)

	//只有两个结构体的所有字段都支持==, !=操作，才可以做比较
	type data1 struct {
		x int
		y map[string]int
	}

	d1 := data1 {
		x : 10,
	}

	d2 := data1 {
		x : 10,
	}

	//fmt.Println(d1 == d2)  ,不可以
	fmt.Println(d1, d2)

	type data2 struct {
		x int
		y string
	}

	d3 := data2 {
		x : 10,
	}

	d4 := data2 {
		x : 10,
	}

	fmt.Println(d3 == d4)      //可以

	//可使用指针直接操作结构体字段(这个很常用),但不能多级指针
	d5 := &data2 {
		x : 10,
		y : "lihui",
	}

	d5.x = 11
	d5.y = "cms"
	fmt.Println(d5)

	//空结构struct{}

	var a struct{}
	var b [10]struct{}
	fmt.Println(a, b)

	//匿名字段:没有名字,仅有类型的字段
	type attr struct {
		perm int
	}

	type file struct {
		name string
		attr              //也叫嵌入字段
	}

	f := &file{
		name : "my.ini",
		attr : attr {     //初始化时必须显示初始化匿名字段
			perm:0755,
		},
	}

	f.perm = 0644
	fmt.Println(f.perm)   //可以直接读取或者设置

    //注意,如果嵌入其他包中类型,隐式字段的名字不带包名
    type data struct {
    	os.File
	}

	d := data {
		File:os.File{},
	}

	fmt.Printf("%#v\n", d)

	//字段标签
	type user struct {
		name string `json:"name"`  //tag
		sex int `json:"sex"`
	}

	u1 := user{"lihui", 1}

	v1 := reflect.ValueOf(u1)
	t1 := v1.Type()

	for i, n := 0, t1.NumField(); i < n; i++ {
		fmt.Printf("%s:%v\n", t1.Field(i).Tag, v1.Field(i))
	}


}