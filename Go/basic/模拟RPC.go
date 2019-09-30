package main

import (
	"fmt"
	"time"
	"github.com/pkg/errors"
)

// RPC（模拟远程过程调用）

/*
服务器开发中会使用RPC（Remote Procedure Call，远程过程调用）
简化进程间通信的过程。RPC 能有效地封装通信过程，
让远程的数据收发通信过程看起来就像本地的函数调用一样。

本例中，使用通道代替 Socket 实现 RPC 的过程。
客户端与服务器运行在同一个进程，
服务器和客户端在两个 goroutine 中运行。
*/

// 模拟RPC客户端的请求和接收消息封装
func RPCClient(ch chan string, req string)(string, error){

	// 向服务器发送请求
	ch <- req

	// 等待服务器返回
	select {
	case ack:= <-ch: // 接收到服务器返回数据
		return ack,nil
	case <-time.After(time.Second):// 超时
		return "", errors.New("Time out")
	}
}

// 模拟RPC服务器端接收客户端请求和回应
func RPCServer(ch chan string){
	for {
		// 接收客户端请求
		data := <-ch

		// 打印接收到的数据
		fmt.Println("server received:", data)

		// 通过睡眠函数让程序执行阻塞2秒的任务
		time.Sleep(time.Second * 2)

		// 反馈给客户端收到
		ch <- "roger"
	}
}

func main() {

	// 创建一个无缓冲字符串通道
	ch := make(chan string)

	// 并发执行服务器逻辑
	go RPCServer(ch)

	// 客户端请求数据和接收数据
	recv, err := RPCClient(ch, "hi")

	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("client received", recv)
	}


}


