package main

import (
	"fmt"
	"time"
)

/**
模拟耗时为1s的任务
*/
func worker(id int, job <-chan int, result chan<- int) {
	for j := range job {
		fmt.Println("worker", id, "processing job")
		time.Sleep(time.Second)
		result <- j * 2
	}

}
func main() {
	//定义job,result为带缓冲的chan，缓冲为100
	job := make(chan int, 100)
	result := make(chan int, 100)

	//启动三个worker
	for w := 1; w <= 3; w++ {
		go worker(w, job, result)
	}

	//发送jobs
	for j := 1; j <= 9; j++ {
		job <- j
	}
	//关闭通道
	close(job)

	//收集返回值
	for a := 1; a <= 9; a++ {
		<-result
	}

}
