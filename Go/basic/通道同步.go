package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	"time"
)

func worker(done chan bool) {

	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("Done")

	done <- true

}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// worker(done)

	// 如果把 <- done 这行代码从程序中移除，程序甚至会在 worker 还没开始运行时就结束了
	<-done
}
