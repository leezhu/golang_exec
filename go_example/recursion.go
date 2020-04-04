//.package go_example

package main

import "fmt"

func face(n int) int {
	if n == 0 {
		return 1
	}
	return n * face(n-1)
}

func main() {
	fmt.Println("face:", face(7))
}
