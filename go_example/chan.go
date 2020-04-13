//.package go_example

//通道
package main

import (
	"fmt"
	"time"
)

func main() {
	//创建size个chan，如果为0，则取出chan会等待
	size := 10
	c1 := make(chan int, size)
	go func() {
		for i := 1; i < 4; i++ {
			val := i * 100
			fmt.Println(time.Now(), "<-", val, "at", i)
			c1 <- val //往通道里塞值
		}
		c1 <- 0
	}()
	time.Sleep(10 * time.Second)
	fmt.Println("after sleep")
	for val := range c1 {
		fmt.Println(time.Now(), " recive:", val)
		if val == 0 {
			break
		}
	}
}
