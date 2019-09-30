package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	"time"
)

// 超时处理
func main() {

	c1 := make(chan string, 1)
	go func () {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()


	// 使用 select 实现一个超时操作
	// 如果这个操作超过了允许的 1 秒 的话，将会执行超时 case
	select {
		case res := <- c1:// res := <- c1 等待结果
			fmt.Println(res)
		case <-time.After(time.Second * 1): //<-Time.After 等待超时 时间 1 秒后发送的值
			fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	//超时时间 3 秒，将会成功的从 c2 接收到值，并且打印出结果
	select{
		case res := <- c2:
			fmt.Println(res)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout 2")
	}


}
