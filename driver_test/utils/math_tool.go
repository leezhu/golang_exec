//.package utils

//数学工具，随机数互斥锁
package utils

import (
	"math/rand"
	"sync"
	"time"
)

var (
	//随机数互斥锁(确保GetRandomInt不能被并发)
	//强制同步
	randomMutex sync.Mutex
)

//获取指定范围内的随机数
func GetRandomInt(start, end int) int {
	randomMutex.Lock()
	<-time.After(1 * time.Nanosecond)                    //可以当作sleep使用
	r := rand.New(rand.NewSource(time.Now().UnixNano())) //当前时间的unixtime作为随机数种子
	n := start + r.Int(end-start+1)
	randomMutex.Unlock()
	return n
}
