//.package go_example

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "bob", age: 20})
	fmt.Println(person{name: "bob"})
	fmt.Println(&person{name: "Bob"})
	s := &person{name: "Sean", age: 50}
	fmt.Println(s.name)
	sp := s
	fmt.Println("sp.name:", sp.name)
}
