//.package go_example

package main

import (
	"fmt"
)

func main() {
	s := make([]string, 3)
	fmt.Println("slice inital:", s)

	//单个赋值
	s[0] = "a"
	s[1] = "b"
	s[2] = "d"
	fmt.Println("set s:", s)

	//append
	s = append(s, "d")
	s = append(s, "c")

	fmt.Println("append s:", s)

	//copy
	c := make([]string, 3)
	copy(c, s)
	fmt.Println("c:", c)

	l := c[2:3]
	fmt.Println("l:", l)

	//slice获取
	l1 := c[2:]
	l2 := c[:3]
	fmt.Println("l1:", l1)
	fmt.Println("l2:", l2)

	//string,mutil,声明并初始化变量
	t := []string{"1", "2", "3"}
	// t[6] = "6"
	fmt.Println("t:", t)

	t = append(t, "12")
	fmt.Println("set t:", t)

	//mutil arr
	twoD := make([][]int, 3) //初始化二位数组的行数
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:=", twoD)

}
