//实例展示如何使用atomic包里面的
//store和load类函数来提供对数值类型
//的安全访问
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	wg sync.WaitGroup
)

func main(){
	wg.Add(2)
	
	go doWork("A")
	go doWork("B")

	time.Sleep(1*time.Second)

	fmt.Println("shutdown Now")
	atomic.StoreInt64(&shutdown,1)
	wg.Wait()
}

//doWork用来模拟执行工作的goroutine.
//检测之前的shutdown标志来决定是否提前终止
func doWork(name string){
	defer wg.Done()

	for{
		fmt.Printf("Doing %s Work\n",name)
		time.Sleep(250*time.Millisecond)

		if atomic.LoadInt64(&shutdown)==1{
			fmt.Printf("shutting %s down\n",name)
			break
		}
	}
}