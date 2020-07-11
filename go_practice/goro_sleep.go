package main

import (
	"fmt"
	"sync"
	// "time"
)

func main() {
	// c:=make(chan struct{})
	// go func(){
	// 	fmt.Println("sleepping...")
	// 	time.Sleep(time.Second*5)
	// 	c<-struct{}{}
	// }()
	// <-c
	// fmt.Println("end")
	funcSlice := make([]func(), 0, 10)
	for i := 0; i < 10; i++ {
		i := i
		f := func() {
			b(i)
		}
		funcSlice = append(funcSlice, f)
	}
	print(funcSlice...)
}

func print(funcSlice ...func()) {
	var wg sync.WaitGroup
	for _, v := range funcSlice {
		v := v
		wg.Add(1)
		go func() {
			v()
			wg.Done()
		}()
	}
	wg.Wait()
}

func b(no int) {
	fmt.Println(no)
}
