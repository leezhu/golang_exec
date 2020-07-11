//等待go协程完成后再退出，用chan无缓冲通道实现同步过程
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}) //struct{}空类型
	defer conn.Close()
	//匿名函数
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} //输入
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done //等待刚才的goroutine完成。
	//因为是无缓冲通道，所以接收者没有等到数据必然阻塞等
}

//将网络返回的数据发送到终端
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
