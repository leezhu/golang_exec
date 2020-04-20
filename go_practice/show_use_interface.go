//.package go_practice

//展示如何使用多态
package main

import (
	"fmt"
)

type nodifier3 interface {
	notify()
}

type user3 struct {
	name  string
	email string
}

func (u *user3) notify() {
	fmt.Printf("Sending name=%s,email=%s", (*u).name, (*u).email)
}

type admin struct {
	name  string
	email string
}

func (u *admin) notify() {
	fmt.Printf("Sending name=%s,email=%s", (*u).name, (*u).email)
}

func sendNotification3(n nodifier3) {
	n.notify()
}

func main() {
	bill := &user3{"me", "qq.com"}
	ad := &admin{"you", "qq.com"}
	sendNotification3(bill)
	sendNotification3(ad)
}
