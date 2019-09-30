package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	_"time"
)

// 使用 jobs 发送 3 个任务到工作函数中，然后 关闭 jobs
func main() {

	// 将使用一个 jobs 通道来传递 main() 中 Go 协程任务执行的结束信息到一个工作 Go 协程中。当我们没有多余的 任务给这个工作 Go 协程时，我们将 close 这个 jobs 通道
	jobs := make(chan int, 5)
	done := make(chan bool)

	//使用 j, more := <- jobs 循环的从 jobs 接收数据。在接收的这个特殊的二值形式的值中， 如果 jobs 已经关闭了，并且通道中所有的值都已经接收 完毕，那么 more 的值将是 false。当我们完成所有 的任务时，将使用这个特性通过 done 通道去进行通知
	go func(){
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job",j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 0; j <=3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
