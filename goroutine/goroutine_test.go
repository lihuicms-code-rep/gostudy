package goroutine

import (
	"fmt"
	"testing"
)

func TestSyncSituation(t *testing.T) {
	fmt.Println("顺序情况函数的调用")
	SyncSituation()
}

func TestASyncSituation(t *testing.T) {
	fmt.Println("goroutine情况函数的调用")
	ASyncSituation()
}

func TestNewClockServer(t *testing.T) {
	fmt.Println("未使用goroutine的clock服务")
	NewClockServer()
}

func TestNewEchoServer(t *testing.T) {
	fmt.Println("并发的echo服务器")
	NewEchoServer()
}

func TestNewClient(t *testing.T) {
	fmt.Println("启动客户端.......")
	NewClient()
}