//.package go_example

//sync一直阻塞主进程，等待子进程完成后再
// Add(1)相当于加一个任务，Done相当于Add(-1),Wait()阻塞
package main

import (
	"fmt"
	"sync"
)

func doSomething(i int, wg *sync.WaitGroup, ch chan int) {
	ch <- i
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1000)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go doSomething(i, &wg, ch)
	}
	wg.Wait() //等待上面的完成
	fmt.Println("all done")
	for i := 0; i < 1000; i++ {
		dd := <-ch
		fmt.Println(dd)
	}
}
