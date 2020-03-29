package main

import "fmt"

//goroutine之间相互通信
//通过channel机制


func worker(c chan int) {
	fmt.Println(" i am worker, get chan value:", <-c)
}


func main() {

	//创建一个chan
	c := make(chan int)

	go worker(c)

	//向c中写入
	c <- 2
	fmt.Println(" i am main")
}
