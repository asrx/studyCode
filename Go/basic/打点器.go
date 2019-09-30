package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	"time"
)

// 打点器 是当你想要在固定的时间间隔重复执行 准备的

/*

*/
func main() {

	// 在这个通道上使用内置的 range 来迭代值每隔 500ms 发送一次的值
	ticker := time.NewTicker(time.Millisecond * 500)
	go func () {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// 在运行后 1600ms 停止这个打点器
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped.")

}
