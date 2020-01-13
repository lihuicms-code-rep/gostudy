package goroutine

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//并发的clock服务,这里以网络编程为例子
func NewEchoServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:12306")
	if err != nil {
		log.Println("listen error ", err)
		return
	}

	for {
		fmt.Println("阻塞等待客户端的连接")
		conn, err := listener.Accept()              //这里就会被阻塞住
		if err != nil {
			log.Println("server accept conn error ", err)
			continue
		} else {
			log.Println("客户端连接过来......")
		}

		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		echo(c, input.Text(), 1 * time.Second)
	}
}

//服务器echo数据到客户端
func echo(c net.Conn, shout string, delay time.Duration) {
    fmt.Fprintln(c, "\t", strings.ToUpper(shout))
    time.Sleep(delay)
    fmt.Fprintln(c, "\t", shout)
    time.Sleep(delay)
    fmt.Fprintln(c,"\t", strings.ToLower(shout))
}

