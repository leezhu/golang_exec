package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	data, err := QuerySalesData(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count=", data, "err=", err)
}

//实现并发请求5个http，主goro设置5s超时，一旦超时需要
//销毁停止子goro避免资源浪费

func QuerySalesData(ctx context.Context) (int, error) {
	parentCtx, _ := context.WithTimeout(ctx, time.Second*5)

	c := make(chan int)
	err := make(chan error) //错误请求返回
	urlCharSlice := []string{"a", "b", "c", "d", "e"}

	for _, v := range urlCharSlice {
		v := v
		go querySales(parentCtx, v, c, err)
	}
	defer close(c)

	sum := 0
	i := 0
	for {
		select {
		case <-err:
			return 0, errors.New("http request wrong")
		case v := <-c:
			sum += v
			i += 1
			if i == 5 {
				fmt.Println("sum=", sum)
				return sum, nil
			}
		case <-parentCtx.Done():
			return 0, parentCtx.Err()
		}
	}
}

type res struct {
	count int
	err   error
}

func querySales(ctx context.Context, s string, c chan<- int, errs chan<- error) {
	//接收父context,请求http服务，失败则写入error chan
	url := fmt.Sprintf("https://interview.moreless.io/questions/async_workers/sales/%s", s)
	dataChan := make(chan *res)

	//当前goroutine销毁后，下面的get请求不会销毁，只能等到http超时自动退出
	go get(url, dataChan)

	var result *res
	select {
	case <-ctx.Done():
		return
	case result = <-dataChan:
		close(dataChan)
		// fmt.Println("result", result)
		if result.err != nil {
			fmt.Println("err====", result.err)
			errs <- result.err
			return
		}
		c <- result.count
		return
	}
}

type Counter struct {
	Count int `json:"count"`
}

func get(url string, r chan<- *res) {
	client := http.Client{Timeout: time.Second * 5}
	resp, err := client.Get(url)
	if err != nil {
		r <- &res{0, err}
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	fmt.Println(string(body), resp.StatusCode)
	var counter Counter
	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &counter)
		if err != nil {
			r <- &res{0, err}
			return
		}
	} else {
		r <- &res{0, errors.New("response code not 200")}
		return
	}
	r <- &res{counter.Count, nil}
	return
}
