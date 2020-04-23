//竞争锁的状态

package main

import(
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
	mutex sync.Mutex
)

func main(){
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("counter end，counter=",counter)

}

func incCounter(count int){
	defer wg.Done()

	for count:=0;count<2;count++{
		//锁住临界区
		mutex.Lock()
		{
			value:=counter
			runtime.Gosched()	//切换协程
			value++
			counter=value
		}
		mutex.Unlock()
		//释放锁，进入临界区
	}
}