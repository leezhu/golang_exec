//.package go_example

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 11
	m["k2"] = 12

	fmt.Println("m:", m)
	fmt.Println("m[\"k1\"]=", m["k1"])

	delete(m, "k1") //删除

	err, prs := m["k2"]
	fmt.Println("prs:", prs, " err:", err) //取可能为空的值

	s := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("s:", s)

	var s2 map[string]int
	s2 = s
	fmt.Println("s2:", s2)

	for k, v := range "go" {
		fmt.Printf("%d->%d \n", k, v)
	}
}
