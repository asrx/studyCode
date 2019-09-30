package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var cond sync.Cond

// 生产者
func producer01(out chan<- int,idx int) {
	for i := 0; i < 10; i++ {
		// 条件变量对应互斥锁加锁
		cond.L.Lock()

		// 产品区满 等待消费者消费 *for*
		for len(out) == 3{
			// 挂起当前协程，等待条件变量满足，被消费者唤醒
			cond.Wait()
		}

		num := rand.Intn(1000)
		out <- num
		fmt.Printf("生产者 %dth，生产：%d\n",idx,num)

		// 生产结束，解锁互斥锁
		cond.L.Unlock()
		// 唤醒 阻塞的 消费者
		cond.Signal()
		// 生产完休息下，给其他协程执行机会
		time.Sleep(time.Millisecond * 500)
	}
	close(out)
}

// 消费者
func consumer01(in <-chan int,idx int)  {

	for {
		cond.L.Lock()

		for len(in) == 0 {
			cond.Wait()
		}

		num := <-in
		fmt.Printf("----消费者 %dth，消费：%d\n",idx,num)

		cond.L.Unlock()
		// 唤醒 阻塞的 生产者
		cond.Signal()

		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())

	// 产品区（公共区）使用channel模拟
	product := make(chan int,3)

	// 创建互斥锁和条件变量
	cond.L = new(sync.Mutex)

	// 生产者
	for i := 0; i < 5; i++ {
		go producer01(product,i+1)
	}

	// 消费者
	for i := 0; i < 3; i++ {
		go consumer01(product,i+1)
	}


	for{
		;
	}

}
