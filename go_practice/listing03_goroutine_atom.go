//竞争状态，会取到副本的脏数据，4次读写，最终应该
//3，但实际是2。改动原子操作，会变成4
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main(){
	wg.Add(2)

	//go
	go inCounter(1)
	go inCounter(2)

	//go: 
	wg.Wait()
	fmt.Println(counter)

	//结果：4
}

func inCounter(id int){
	defer wg.Done()
	for count:=0;count<2;count++{
		//安全的加1
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
}