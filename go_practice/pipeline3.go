//.package

package main

import "fmt"

//计数，放入counter中
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

//squarer
func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

//printer
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
