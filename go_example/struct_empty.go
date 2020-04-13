//.package go_example
//练习struct{}复合类型。

package main

import "fmt"

func main() {
	var set map[string]struct{}
	set = make(map[string]struct{})
	fmt.Println("begin:", set)
	set["first"] = struct{}{} //可以把strcut{}想成 type my_struct struct{}
	//那么这句就可以表达成 set["first"]=my_struct{""},string映射的是一个struct类型
	set["seconde"] = struct{}{}
	fmt.Println("end:", set)
	// var b []int
	// b = append(b, 1)
	// var c map[string]string
	// c["wo"] = "ni"
	// fmt.Println(c["wo"])
	// fmt.Println()
	// a := make([]int, 0, 0)
	// fmt.Printf("b:%v,len(b)=%d,cap(b)=%d", b, len(b), cap(b))
	// fmt.Println(append(b, 1, 2, 3, 4))
	//初始化
	// set = make(map[string]struct{})
}
