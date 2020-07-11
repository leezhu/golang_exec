//并发回声服务器
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConn(con)
	}
}

//回声
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//处理器
func handleConn(con net.Conn) {
	input := bufio.NewScanner(con) //接受con
	for input.Scan() {
		go echo(con, input.Text(), 1*time.Second) //接收客户端传的信息
	}
	con.Close()
}
