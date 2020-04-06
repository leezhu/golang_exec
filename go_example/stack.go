//.package go_example
// 在栈的题时，想到go是没有栈的数据类型的，可以用slice进行模拟
package main

import "fmt"
import "errors"

var stack []int //先进后出,全局分片

func push(value int) {
	stack = append(stack, value)
}

//从栈中获取数据
func pop() (int, error) {
	value := 0
	if len(stack) > 0 {
		value = stack[len(stack)-1]
		stack = stack[:len(stack)-1] //这里是关键，如果长度为1时,stack=[:0]取到为空
		return value, nil
	} else {
		return 0, errors.New("stack is empty")
	}
}

func main() {
	//添加数据
	for i := 0; i < 10; i++ {
		push(i)
		fmt.Println(stack)
	}
	//从栈取出数据
	for {
		v, err := pop()
		if err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println("v=", v)
		}
	}
}

func removeOuterParentheses(S string) string {
	level := 0
	var res []rune
	for _, v := range S {
		if v == 41 {
			level--
		}
		if level >= 1 {
			res = append(res, v)
		}
		if v == 40 {
			level++
		}
	}
    return string(res)
}
