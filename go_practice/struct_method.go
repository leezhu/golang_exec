//.package go_practice

//结构体方法
package main

import "fmt"

type new_user struct {
	name  string
	email string
}

//值类型不可改变结构体原值
func (u new_user) notify(name string) {
	fmt.Printf("send user email to %s %s", u.name, u.email)
	u.name = name
}

//指针类型可以改变原值
func (u *new_user) changeEmail(email string) {
	u.email = email
}

func main() {
	var user new_user
	fmt.Println("username1:", user.name)
	user.notify("yes")
	fmt.Println("username2:", user.name)

	fmt.Println("useremail1:", user.email)
	user.changeEmail("new@qq.com")
	fmt.Println("useremail2:", user.email)
}
