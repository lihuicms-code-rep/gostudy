package main

import (
	"fmt"
)

//无序键值对集合,key必须是支持==,!=运算的类型, 数字,指针,字符串,指针，数组，结构体
func main() {
	//m := make(map[string]int)
	//fmt.Printf("m type:%T, addr:%p, value:%v, len:%d\n", m, m, m, len(m))
	//m["a"] = 1
	//m["b"] = 1
	//
	////value部分是匿名结构体
	//m1 := map[int]struct{x int}{
	//	1 : {1},
	//	2 : {2},
	//	3 : {3},
	//}
	//
	//fmt.Println(m1)
	//
	////ok-idiom写法
	//if v, ok := m["c"]; ok {
	//	fmt.Printf("c的值:%d\n", v)
	//} else {
	//	fmt.Println("没有c这个key")
	//}
	//
	////对字典进行迭代,次序都不同
	//
	////对map value部分的修改
	//m["a"] += 2
	////但是如果value部分是结构体或者数组的,不可以修改value成员
	////考虑到内存访问安全和哈希算法等原因,not addressable
	////m1[1].x += 1
	//
	////那这种情况如何修改呢？
	////这个也是在自己开发中遇到的情景
	////1.整个value部分修改,然后设置
	//v := m1[1]
	//v.x = 2
	//m1[1] = v
	//fmt.Println("update m1", m1)
	//
	////2.value部分设计成指针类型
	//type user struct {
	//	name string
	//	age int
	//}
	//
	//m2 := map[int]*user{
	//	1:&user{
	//		name : "lihui",
	//		age : 11,
	//	},
	//}
	//
	//m2[1].age += 1
	//
	//fmt.Println(m2)
	//
	//
	////未初始化的map,不同于通过make的,特别要注意初始化的作用
	//var m3 map[string]int
	//fmt.Println(m3 == nil)   //true
	//fmt.Printf("m3 addr:%p, value:%+v\n", m3, m3)
	//
    //var m4 = map[string]int{}
    //fmt.Println(m4 == nil)   //false
    //fmt.Printf("m4 addr:%p, value:%+v\n", m4, m4)
	//
	//
    ////安全
    ////在迭代期间删除或者新增键值是安全的
    //m5 := make(map[int]int)
    //for i := 0; i < 10; i++ {
    //	m5[i] = i
	//}
	//
    //for k := range m5 {
    //	if k == 5 {
    //		m5[k] = 50000000000
    //		delete(m5, 6)
    //		m5[20] = 100
	//	}
    //	fmt.Println(k, "===>", m5)
	//}



    //对map的数据竞争
    //运行时会对字典并发操作做出检测,如果某个任务正在对字典进行写操作，那么其他任务就不可对该字典读写操作
    //m6 := make(map[string]int)
	//
    //go func() {
    //	for {
    //		m6["a"] += 1
    //		time.Sleep(time.Microsecond)
	//	}
	//}()
	//
    //go func() {
    //	for {
    //		_ = m6["b"]
    //		time.Sleep(time.Microsecond)
	//	}
	//}()
	//
	//select {}    //阻塞

	//使用sync.RWMutex来实现同步
	//var lock sync.RWMutex
	//m7 := make(map[string]int)
	//
	//go func() {
	//	for {
	//		lock.Lock()
	//		m7["a"] += 1
	//		lock.Unlock()
	//		time.Sleep(time.Microsecond)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		lock.RLock()
	//		_ = m7["b"]
	//		lock.RUnlock()
	//		time.Sleep(time.Microsecond)
	//	}
	//}()


	//性能,字典对象本身就是指针包装，传参时无须再取地址

	//在创建时预先准备好空间有助于提升性能,减少扩张时内存分配和重新哈希操作
	m8 := make(map[string]int)
	m9 := make(map[string]int, 10)

	fmt.Println(len(m8), m8)
	fmt.Println(len(m9), m9)


	//select {}    //阻塞
}

