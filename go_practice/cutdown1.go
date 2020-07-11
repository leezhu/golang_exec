package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)

	go func() {
		for countdown := 10; countdown > 0; countdown-- {
			fmt.Println(countdown)
			<-tick
		}
	}()

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	fmt.Println("commencding countdown,press return concel")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("lanche concel")
		return
	}
	fmt.Println("lanch")
}
