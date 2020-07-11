//.package
//两个channal连接数据
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x //执行平方后放入squares
		}
		close(squares)
	}()

	//printer输出
	for x := range squares {
		fmt.Println(x)
	}
}
