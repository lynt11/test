package main

import (
	"fmt"
	"time"
)

func main() {
	//接收请求的通道
	request := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		request <- i
	}
	close(request)

	//limiter通道每200ms接收一个值，也是速率限制中的管理器
	limiter := time.Tick(time.Millisecond * 200)
	for req := range request {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	//通道缓冲
	burstyLimiter := make(chan time.Time, 3)
	//将通道缓冲设置三次临时的值
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	/**
	每200ms添加一个新的值进入buerstyLimiter
	直到到达3个限制
	*/
	go func() {
		for t := range time.Tick(time.Microsecond * 200) {
			burstyLimiter <- t
		}
	}()

	/**
	模拟超过五个的请求，刚开始的三个将受burstyLimiter的脉冲影响
	*/
	burstyRequest := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for req := range burstyRequest {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

}
