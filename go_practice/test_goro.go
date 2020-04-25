//检测协程运行时间长，主进程main是否会退出
package main

import (
	"fmt"
	"time"
	"sync"
)
var sy sync.WaitGroup
func run(){
	defer sy.Done()
	fmt.Println("开始协程任务")
	time.Sleep(time.Second*10)
	fmt.Println("结束协程任务")
}
func main(){
	sy.Add(1)
	fmt.Println("开始主进程任务")
	go run()
	sy.Wait()
	fmt.Println("结束主进程任务")
}