package goroutine

import (
	"fmt"
	"time"
)

//go routine的简单举例
func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func Fib(x int) int {
	if x < 2 {
		return x
	}

	return Fib(x-1) + Fib(x-2)
}

//这种写法的话,会一直阻塞在Spinner上
func SyncSituation() {
	Spinner(100 * time.Millisecond)
	fibN := Fib(10)
	fmt.Printf("Fib(%d)=%d", 10, fibN)
}

//goroutine的使用
func ASyncSituation() {
	go Spinner(100 * time.Millisecond)
	fibN := Fib(50)    //这是一个比较耗时的行为
	fmt.Printf("Fib(%d)=%d", 50, fibN)
}