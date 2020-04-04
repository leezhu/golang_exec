//.package go_example

package main

import "fmt"

type geometry interface {
	area() float64
	// prim() float64
}

type rect struct {
	width, height float64
}

type cel struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height * 1.0 / 2
}
func (c cel) area() float64 {
	return c.width * c.height * 3.14
}

func measure(g geometry) {
	fmt.Println("area=", g.area())
}

func main() {
	r := rect{width: 12, height: 23}
	c := cel{23, 12}
	measure(r)
	measure(c)
}
