//协程练习，chanal

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	//分配一个模拟器，观察两个协程在一个处理器上的并发
	runtime.GOMAXPROCS(2)

	//wg用来等待程序完成
	//计数加2，表示要等待两个goroutine结束
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutine")

	//声明一个匿名函数，并创建一个goroutine
	//小写字母
	go func(){
		//在函数退出时调用Done代表结束
		defer wg.Done()

		//显示字母表3次
		for count:=0;count<3;count++{
			for char:='a';char<'a'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()


	//声明一个匿名函数，并创建一个goroutine
	//大写字母
	go func(){
		//在函数退出时调用Done代表结束
		defer wg.Done()

		//显示字母表3次
		for count:=0;count<3;count++{
			for char:='A';char<'A'+26;char++{
				fmt.Printf("%c",char)
			}
		}
	}()
	fmt.Println("Waiting To Finish")
	wg.Wait()
	fmt.Println("\nTermimalting program")
}

