package main

import "fmt"

func f3() {
	fmt.Println("f")
}

func f1(f func()) {
	fmt.Println("begin")
	f()
	fmt.Println("end")
}

func main() {
	f1(f3)
}
