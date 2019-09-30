package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0

	fmt.Println(s)

	for _, v := range s{
		sum += v
	}

	c <- sum // 把 sum 发送到通道 c
}

func main() {
	
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	fmt.Println(len(s)/2)

	go sum(s[:len(s)/2] , c)

	// x := <-c

	go sum(s[len(s)/2:], c)

	// y := <-c

	x, y := <-c, <-c // 从通道c中接收

	fmt.Println("x:", x, ",y:", y, ",x+y:", x+y)

	// fmt.Println("x:",x)
	// fmt.Println("y:",y)

}


