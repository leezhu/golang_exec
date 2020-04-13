//.package craw_spider_coro

//并发爬虫,整体设计是先等到获取指定量的任务的图片下载链接数量，
//然后进行下载

package main

import (
	"fmt"
	"sync"
)

var (
	//存放图片链接
	chanImgUrls chan string //存放string类型的管道
	//存放任务是否完成
	chanTask  chan string
	waitGroup sync.WaitGroup //任务完成
)

//获取一个页面上所有的链接
func SpiderPrettyImg(url string) (urls []string) {
	// pageStr := GetPageStr(url)
	return urls
}

//图片处理
func SpiderImgUrls(url string) {
	//获取一个页面下的所有图片
	urls := SpiderPrettyImg(url)
	for _, url := range urls {
		chanImgUrls <- url //放入url通道中
	}

	//通知协程任务完成
	chanTask <- url  //url下载任务中
	waitGroup.Done() //说明已完成
}

func CheckIfAllSpidersOk() {
	var count int
	for {
		url := <-chanTask //阻塞等待任务
		fmt.Printf("%s完成爬取任务\n", url)
		count++
		if count == 147 { //如果完成了147个任务，那么关闭图片的通道存入
			close(chanImgUrls)
			break
		}
	}
	waitGroup.Done()
}

func main() {
	//初始化数据管道
	chanImgUrls = make(chan string, 100000)
	chanTask = make(chan string, 147)

	//爬虫协程：源源不断的往通道里中添加图片链接
	for i := 1; i < 148; i++ {
		waitGroup.Add(1)
		go SpiderImgUrls("https://hbimg.huabanimg.com/0a6efb72646cc890c0f8e5ca3b506847842c703d91584-Guqow8_fw658")
	}

	//开辟任务统计协程，如果都完成，则关闭数据通道
	waitGroup.Add(1)
	go CheckIfAllSpidersOk()

	//下载协程：源源不断地从管道中获取地址并下载
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go DownloadImg()
	}
	waitGroup.Wait()
}
