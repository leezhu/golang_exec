//.package utils
//错误处理工作

package main

import (
	"fmt"
	"os"
)

//处理错误，有错误时暴力退出
func HandlerError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}
