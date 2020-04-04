//协程
package main

import "fmt"

func f(from string) {
	for i := 0; i < 30; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	go func(msg string) {
		fmt.Println(msg)
	}("going") //匿名函数，传值

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
