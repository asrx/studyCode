package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	"time"
)

// 工作池
func worker(id int, jobs <-chan int, results chan<- int) {
	// 从 jobs 通道接收任务
	// 通过 results 发送对应 的结果
	for j:= range jobs{
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)	//让每个任务间隔 1s 来模仿一个耗时的任务
		results <- j * 2
	}
}

/*

*/
func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w<=3; w++{
		// 启动了 3 个 worker，初始是阻塞的，因为 还没有传递任务
		go worker(w,jobs,results)
	}

	// 发送 9 个 jobs，然后 close 这些通道 来表示这些就是所有的任务了
	for j := 0; j < 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}

}
