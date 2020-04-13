//.package go_example

//用通道实现监督和通过简单的代码实现，一个协程实现具体的工作。
//另外一个协程用户检查程序是否消费完成。
package main

import (
	"fmt"
	"time"
)

func main() {
	workers := 3
	ch := make(chan int) //创建通道
	// worker :=
	// leader :=
	go func() {
		cnt := 0 //在协程内部
		fmt.Println("监督协程")
		//监督协程，看是否完成
		for data := range ch { //只看有多少数量
			// a := <-ch
			fmt.Println(data)
			cnt++
			if cnt == workers {
				fmt.Println("已完成")
				break
			}
		}
		close(ch)
	}()
	fmt.Println("协程呢")
	for i := 0; i < workers; i++ {
		fmt.Printf("循环了第%d次\n", i)
		go func(i int) {
			fmt.Println("放入协程")
			time.Sleep(2 * time.Second)
			ch <- i
		}(i) //工作协程，做完后将数据写入通道中
	}
}
