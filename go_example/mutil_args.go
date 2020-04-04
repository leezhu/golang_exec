//.package go_example

package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for num := range nums {
		total += num
	}
	fmt.Println("total=", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3, 4)
	nums := []int{1, 2, 3, 4, 5, 5, 676, 9}
	sum(nums...)
}
