package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for j := 1; j < 10; j++ {
		fmt.Println(j)
	}
}
