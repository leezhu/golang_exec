//竞争状态，会取到副本的脏数据，4次读写，最终应该
//3，但实际是2
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
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

	//结果：2
}

func inCounter(id int){
	defer wg.Done()
	for count:=0;count<2;count++{
		value:=counter
		runtime.Gosched()
		value++
		counter=value
	}
}