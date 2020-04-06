//.package go_example

//前面stack.go做了简单的栈，但是因为会有stack整个复制位移操作，效率不高。
//现在参照大佬代码实现一个效率高的版本
package main

import "fmt"

type stack struct {
	s   []byte //int8，存储字符
	top int    //末尾标记位
}

//推入
func (s *stack) Push(c byte) {
	s.s = append(s.s, c)
	s.top++ //第一个元素top=1
}

//pop弹出
func (s *stack) Pop() byte {
	s.top--
	c := s.s[s.top] //获取第一个元素s.top=0
	s.s = s.s[:s.top]
	return c
}

//为空判断
func (s *stack) Empty() bool {
	return s.top == 0
}

//顶端top数据
func (s *stack) Top() byte {
	if !s.Empty() {
		return s.s[s.top-1]
	}
	return -1

}

func main() {
	var s stack
	bs := []byte(s) //将string转成byte数组，range遍历也是可以得到字符
	for i := 0; i < len(bs); i++ {
		tmp := s.Top()
		if tmp == bs[i] {
			s.Pop()
		} else {
			s.Push(bs[i])
		}
	}
	fmt.Println(string(stack)) //string可以将 rune或者byte数组转为字符串。
}
