//.package go_example

package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, 4, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	s := map[string]int{"a": 1, "b": 2}
	var b map[string]int
	b = make(map[string]int)
	b["a"] = 1
	fmt.Println("s:", s, "b:", b)
	for k, v := range s {
		fmt.Println("k=", k, "v=", v)
	}
}
