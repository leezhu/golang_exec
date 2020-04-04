//.package go_example

package main

import (
	"fmt"
)

func closure() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextIn := closure()
	fmt.Println("nextIn:", nextIn())
	fmt.Println("nextIn:", nextIn())
	fmt.Println("nextIn:", nextIn())
	fmt.Println("nextIn:", nextIn())
}
