//.package go_practice

//buffer.reader
package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var bar bytes.Buffer

	//将字符串写入buffer
	var bytes_data = []byte("me")
	bar.Write(bytes_data)

	//使用Fprintf将字符串拼接到buffer
	fmt.Fprintf(&bar, " love")
	io.Copy(os.Stdout, &bar)
}
