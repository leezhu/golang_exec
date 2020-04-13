//.package go_practice

//curl监听，简单版本的curl实现

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//init函数在main函数之前运行
func init() {
	if len(os.Args) != 2 {
		fmt.Println("usage..")
		os.Exit(-1)
	}
}

func main() {
	fmt.Println(os.Args[1])
	r, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println("get error:", err)
		return
	}

	//从Body复制到Stout,标准输出
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println("read body error:", err)
	}
}
