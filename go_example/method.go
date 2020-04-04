//.package go_example

package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) prim() int {
	return r.width / r.height
}

func main() {
	s := rect{width: 1, height: 2}
	fmt.Println("area=", s.area())
	fmt.Println("prim=", s.prim())
	b := s
	b.width = 20
	fmt.Println("change,area=", s.area())
	fmt.Println("change,b.area=", b.area())
	fmt.Println(s == b)

	ptr := &s
	fmt.Println("ptr,area=", ptr.area())
}
