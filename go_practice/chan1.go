//.package,缓冲通道

package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	ch <- "C"
	fmt.Println(<-ch)
	fmt.Println(cap(ch))
	fmt.Println(len(ch))
}
