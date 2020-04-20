//.package go_practice

//嵌入式方法，新类型的方法在原有方法上的扩展
package main

import "fmt"

type user struct {
	name  string
	email string
}

//notify实现了一个方法
func (u *user) notify() {
	fmt.Printf("Seding,name=%s,email=%s", u.name, u.email)
}

//admin代表一个拥有权限的管理人员
type admin struct {
	*user //嵌入类型
	level string
}

func main() {
	//初始化结构体，key不是字符
	//内部类型的初始化是用结构字面量
	ad := &admin{
		user: &user{
			name:  "you",
			email: "11.com",
		},
		level: "1",
	}
	//嵌入方法提升到外部
	ad.notify()

	//也可以使用嵌入方法
	ad.user.notify()
}
