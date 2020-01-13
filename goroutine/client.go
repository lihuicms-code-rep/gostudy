package goroutine

import (
	"io"
	"log"
	"net"
	"os"
)

func NewClient() {
	conn, err := net.Dial("tcp", "localhost:12306")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	go mustCopy(os.Stdout, conn)     //从conn拷贝来自服务器的数据至标准输出
	mustCopy(conn, os.Stdin)         //从标准输入拷贝数据至conn
}

//从源reader读取数据至目标writer
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}