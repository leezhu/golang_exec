//.package go_practice

//检测方法集的使用,接口一般定义的是一个值的方法
//实现接口的struct 方法如果是指针接受方法
//那么在接口函数传参的时候，需要传入指针类型。
//其实就是struct类型自动转interface实现多态时
//需要注意的问题

package main

import (
	"fmt"
)

//notifier2接口定义了一个notify方法
type notifier2 interface {
	notify()
}

//user结构实现这个接口方法
type user2 struct {
	name  string
	email string
}

//实现指针接受方法
func (u user2) notify() {
	fmt.Println("user2 notify")
}

//nodifier2接口接受
func sendNotification2(n notifier2) {
	n.notify()
}

func main() {
	bill := &user2{"me", "qq.com"}
	sendNotification2(bill)
}
