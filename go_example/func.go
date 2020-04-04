//.package go_example

package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plus_mutil(a, b, c int) (bool, int) {
	return true, a + b + c
}

func main() {
	// res := plus(1, 2)
	// fmt.Println("1+2:", res)
	ok, mutil_data := plus_mutil(1, 2, 3)
	fmt.Printf("bool=", ok, "1+2+3=", mutil_data)
}
