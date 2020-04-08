//.package go_practice

//练习用户自定义类型struct
package main

import "fmt"

type user struct {
	name       string
	email      string
	ext        string
	privileged int
}

type admin struct {
	person user
	level  int
}

type Duration int64

func main() {
	//---------用户自定义的struct类型-------
	// fred := admin{ //多重struct嵌套初始化赋值
	// 	person: user{
	// 		name:       "Micle",
	// 		email:      "yes@qq.com",
	// 		ext:        "mp3",
	// 		privileged: 1,
	// 	},
	// 	level: 1,
	// }
	// fmt.Println(fred)
	//-----------------------

	//-------内置类型的升级版本-----------------
	// var du Duration	//属于不同的类型
	// du = int64(10000) //cannot use int64(10000) (type int64) as type Duration in assignment
}
