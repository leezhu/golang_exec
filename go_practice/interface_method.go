//.package go_practice

//展示go语言如何使用接口
package main

import "fmt"

//notifier是一个定义了
//通知类行为的接口
type notifier interface {
	notify()
}

//user在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

//notify是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("sending user email to %s<%s>", u.name, u.email)
}

//main 是应用程序的入口
func main() {
	//创建一个user类型的值，并发送通知
	u := user{"Bill", "bill@email.com"}

	sendNotification(u)
}

func sendNotification(n notifier) {
	n.notify()
}
