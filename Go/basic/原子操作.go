package main

import (
	"sync/atomic"
	"fmt"
)

// 本例中只是对变量进行增减操作，
// 虽然可以使用互斥锁（sync.Mutex）解决竞态问题，
// 但是对性能消耗较大。在这种情况下，
// 推荐使用原子操作（atomic）进行变量操作

var seq int64

func GenID() int64{
	return atomic.AddInt64(&seq, 1)
	//return seq
}

func main() {

	for i := 0; i < 10; i++ {
		go GenID()
	}


	fmt.Println(GenID())
}



