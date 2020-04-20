//.package go_practice

//展示嵌入式如何运用于接口
//嵌入式的是一个实现接口的struct
//看嵌入式的struct能否用于多态
package main

import (
	"fmt"
)

type notifier4 interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user,name=%s,email=%s", u.name, u.email)
}

type admin struct {
	user
	level string
}

func sendNotification3(n notifier4) {
	n.notify()
}

func main() {
	ad := admin{
		user: user{
			name:  "you",
			email: "11.com",
		},
		level: "1",
	}
	sendNotification3(&ad)
}