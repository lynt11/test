package main


import (
	"errors"
	"fmt"
	"time"
)

//主函数
func main(){
	//创建一个无缓冲字符通道
	ch:=make(chan string)
	//并发执行服务器逻辑
	go RPCSever(ch)
	//客户端请求和接收数据
	recv,err :=RPCClient(ch,"hi")
	if err !=nil{
		//发生错误
		fmt.Println("err")
	} else {
		//正常接收数据
		fmt.Println("client received",recv)
	}
}
//模拟RPC客户端的请求和接受消息封装
func RPCClient(ch chan string,req string)(string,error){
	ch<-req
	select {
	case ack:=<-ch: //接收到服务器返回数据
		return ack,nil
	case <-time.After(time.Second)://超时
	return "",errors.New("time out")

	}
}
//服务器接收和反馈数据
func RPCSever(ch chan string) {
	for  {
		//接收客户端请求
		data:=<-ch
		//打印接收到的数据
		fmt.Println("sever received:",data)
		//模拟超时
		//time.Sleep(time.Second*2)
		//向客户端反馈已收到
		ch<-"roger"
	}
}

