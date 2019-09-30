package main

import (
	"math/rand"
	"time"
	"fmt"
	"sync"
)

var rwMutex02 sync.RWMutex	// 锁只有一把，2个属性

var value int				// 定义全局变量，模拟共享数据

func readGo02(idx int)  {
	for {

		rwMutex02.RLock() 	// 以读模式加锁
		num := value
		fmt.Printf("---%dth 读go程，读出：%d\n",idx,num)

		rwMutex02.RUnlock()	// 以读模式解锁
	}

}

func writeGo02(idx int)  {
	for {
		// 生成随机数
		num := rand.Intn(1000)

		rwMutex02.Lock()	// 以写模式加锁
		value = num
		fmt.Printf("%dth 写go程，写入：%d\n",idx,num)
		time.Sleep(time.Millisecond * 300)// 放大实现现象

		rwMutex02.Unlock()
	}
}

func main() {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	//quit := make(chan bool)	// 用于 关闭主go程的channel
	//ch := make(chan int)	// 用于 数据传递的channel

	for i := 0; i < 5; i++ {
		go readGo02(i+1)
	}

	for i := 0; i < 5; i++ {
		go writeGo02(i+1)
	}

	//<-quit
	for {
		;
	}
}
