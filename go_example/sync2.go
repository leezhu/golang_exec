//.package go_example

//sync一直阻塞主进程，等待子进程完成后再
// Add(1)相当于加一个任务，Done相当于Add(-1),Wait()阻塞
package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	workers := 3
	wg.Add(workers)
	worker := func() {
		defer wg.Done()
		//做一些
		// fmt.Println("begin")
	}
	leader := func() {
		wg.Wait()
	}
	go leader()
	for i := 0; i < workers; i++ {
		go worker()
	}
}
