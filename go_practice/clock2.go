//实现一个传入时间地域和端口的时钟
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "端口?")
	flag.Parse()
	url := fmt.Sprintf("%s:%d", "localhost", port)
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Println(err.Error())
	}
	for {
		conn, _ := listener.Accept()
		handlerConn(conn, port)
	}

}
func handlerConn(c net.Conn, port int) {
	defer c.Close()
	timeFormat := "2006-01-02 15:04:05 Mon\n"
	// now := time.Now()
	var localtion *time.Location
	switch port {
	case 8010:
		localtion, _ = time.LoadLocation("US/Eastern")
		break
	case 8030:
		localtion, _ = time.LoadLocation("Asia/Tokyo")
		break
	case 8020:
		localtion, _ = time.LoadLocation("Europe/London")
		break
	default:
		io.WriteString(c, "-p 8080 ")
		break
	}
	for {
		_, err := io.WriteString(c, time.Now().In(localtion).Format(timeFormat))
		if err != nil {
			log.Fatal(err.Error())
			return
		}
		time.Sleep(1 * time.Second)
	}
}
