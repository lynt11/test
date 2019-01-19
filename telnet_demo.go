package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//telnet回音服务器

//接受连接
//会话处理
//Telnet命令处理
//程序入口
//-------------------------------------------------------
//服务器端
//传入地址和退出的通道
func server(address string, exitChan chan int) {
	//根据给定地址进行侦听
	l, err := net.Listen("tcp", address)
	//处理异常
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	//如果成功打印侦听地址
	fmt.Println("listmen:" + address)
	//延迟关闭侦听器
	defer l.Close()
	//侦听循环
	for {
		//新连接没有到来时，Accept是阻塞的
		conn, err := l.Accept()
		//发生任何的侦听错误，打印错误并推出服务器
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		//根据连接开启会话，这个过程需要并行执行
		go handleSession(conn, exitChan)
	}
}

//------------------------------------------------------------

//连接的会话逻辑
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session start ed:")
	//创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)
	//接收数据的循环
	for {
		//读取字符串，直到碰到回车返回个
		str, err := reader.ReadString('\n')
		//数据读取正确
		if err == nil {
			//去掉字符串尾部的回车
			str = strings.TrimSpace(str)
			//处理telnet命令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			//echo逻辑,发什么数据，原样返回
			conn.Write([]byte(str + "\r\n"))
		} else {
			//发生错误
			fmt.Println("Session closed")
			conn.Close()
			break
		}

	}
}

//telnet命令处理
func processTelnetCommand(str string, exitChan chan int) bool {
	//@close指令表示终止本次对话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		//告诉外部需要断开连接
		return false
		//shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Sever shutdown")
		//往通道中写入0，阻塞等待接收方处理
		exitChan <- 0
		//告诉外部需要断开连接
		return false
	}
	//打印输入的字符串
	fmt.Println(str)
	return true
}

//主流程
func main() {
	//创建一个程序结束码的通道
	exitchan := make(chan int)
	//将服务器并发运行
	go server("127.0.0.1:7001", exitchan)
	//通道阻塞，等待接收返回值
	code := <-exitchan
	//标记程序返回值并退出
	os.Exit(code)
}
