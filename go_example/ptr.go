//.package go_example

package main

import "fmt"

//传值
func zeroval(val int) {
	val = 0
}

//传指针
func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("intial=", i)
	zeroval(i)
	fmt.Println("zeroval=", i)
	zeroptr(&i)
	fmt.Println("zeroptr=", i)
	fmt.Println("&i=", &i)
}
