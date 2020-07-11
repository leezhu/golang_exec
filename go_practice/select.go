package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("begin select")
	c := make(chan string)
	go func() {
		d := bufio.NewScanner(os.Stdin)
		for d.Scan() {
			c <- d.Text()
		}
	}()
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("sleep after")
	case <-c:
		fmt.Println("concel")
		close(c)
	}
	fmt.Println("end select")
}
