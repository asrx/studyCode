package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)		// 用来进行数据通信的 channel
	quit := make(chan bool)		// 用来判断是否退出的 channel

	go func() {				// 写数据

		for{
			select {
			case v := <-ch:
				fmt.Println("num:",v)

			case <-time.After(3 * time.Second):
				fmt.Println("timeout")
				quit<-true
				//break
				// return
				//runtime.Goexit()

				goto label
			}
		}

	label:
		fmt.Println("break to label ----")
	}()

	for i := 0; i < 2; i++ {
		ch <-i
		time.Sleep(time.Second * 2)
	}

	<-quit

	fmt.Println("finsh")
}
