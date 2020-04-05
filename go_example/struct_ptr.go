//.package ptr
//测试函数返回结构体指针或者值有无区别

package main

import (
	"fmt"
)

type Test struct {
	age int
}

type ptr interface {
	GetString(string)
}

func get_struct(age int) (Test, error) {
	return Test{age: 20}, nil
}

func main() {
	new_struct, _ := get_struct(100)
	new_struct.age = 200
	fmt.Println("new_struct:", new_struct)

	new2_struct, _ := get_struct(100)
	// new2_struct.age = 300
	fmt.Println("new2_struct:", new2_struct)
}

//从结果来看，返回的结构体指针和结构体都能够改变原来的值
