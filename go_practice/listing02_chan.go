//实践并发交替执行的情况

package main

import (
	"fmt"
	"runtime"
	"sync"
)

//wg 用来等待程序完成
var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(1)

	wg.Add(2)
	go printPrime("A")
	go printPrime("B")
	wg.Wait()
	fmt.Println("Create Gorountine")
}

func printPrime(prefix string){
	defer wg.Done()
next:
	for outer:=2;outer<5000;outer++{
		for inner:=2;inner<outer;inner++{
			if outer%inner==0{
				continue next
				// continue
			}
		}
		fmt.Printf("%s:%d\n",prefix,outer)
	}
	fmt.Println("Completed",prefix)
}