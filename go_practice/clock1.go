//.package
package main

import (
	"io"
	"net"
	"time"

	"github.com/gpmgo/gopm/modules/log"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(log.DEBUG, err.Error())
			continue
		}
		handleConn(conn)
	}
}

//handleConn 接受清楚处理器
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return //断开
		}
		time.Sleep(1 * time.Second)
	}
}
