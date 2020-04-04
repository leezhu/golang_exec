// package.package go_example

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len(a):", len(a))

	b := [6]int{1, 2, 4, 5, 6}
	fmt.Println("b:", b)

	b[4] = 100
	fmt.Println("set b:", b)
	fmt.Println("get:", b)

	var twoArr [10][10]int
	fmt.Println("towArr begin:", twoArr)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			twoArr[i][j] = i * j
		}
	}
	fmt.Println("towArr set:", twoArr)
}
