package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool) {

	for{
		select {
		case num:=<-ch:
			fmt.Print(num," ")

		case <-quit:
			runtime.Goexit()
		}
	}

}

func main() {

	ch := make(chan int)		// 用来进行数据通信的 channel
	quit := make(chan bool)		// 用来判断是否退出的 channel

	go fibonacci(ch,quit)	// 打印 斐波那契数列

	x, y := 1, 1
	for i := 0; i < 20; i++ {
		ch <- x
		x,y = y, x+y
	}

	quit <- true
}
