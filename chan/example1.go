package main

import (
	"fmt"
	"time"
)

func goWorker(name string) {
	for i := 0; i < 10; i ++ {
		fmt.Println("我是一个go routine, name is ", name)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("go routine %s finished\n", name)
}

func main() {
	//开辟一个go去执行
	//含义:不再让main去执行,让另外的去执行
	go goWorker("小黑")
	go goWorker("小白")

	for i := 0; i < 10; i++ {
		fmt.Println("我是main")
		time.Sleep(time.Second)
	}
}
