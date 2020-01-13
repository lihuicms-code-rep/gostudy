package goroutine

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

//并发的clock服务,这里以网络编程为例子
func NewClockServer() {
	listener, err := net.Listen("tcp", "0.0.0.0:12306")
	if err != nil {
		log.Println("listen error ", err)
		return
	}

	for {
		fmt.Println("一直阻塞等待客户端的连接")
		conn, err := listener.Accept()              //这里就会被阻塞住
		if err != nil {
			log.Println("server accept conn error ", err)
			continue
		} else {
			log.Println("客户端连接过来......")
		}

		go HandleConn(conn)
		//HanderConnSimple(conn)
	}

	fmt.Println("can XXXXXXXXXXXXXXXXX")
	select {   //阻塞住不要退出

	}
}



//连接处理函数,这里每隔2秒向客户端写一个当前服务器时间
//具体逻辑是用for死循环,没有return情况发生就一直阻塞在这里
func HandleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}

		time.Sleep(2 * time.Second)
	}
}

//假如这里不阻塞,做完就OK
func HanderConnSimple(c net.Conn) {
	fmt.Println("已经处理客户端连接")
	_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
	if err != nil {
		return
	}
}


