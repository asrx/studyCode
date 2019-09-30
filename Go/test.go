package main

import (
    "fmt"
)

func main() {

	arr := []int{1,2,3}

	arr2 := append(arr,4)

	fmt.Println(arr2)

	// ch1 := make(chan int, 1)
	// ch2 := make(chan int, 1)
	// ch3 := make(chan int, 1)

	// go func() {
	// 	fmt.Println("1")
	// 	ch1 <- 1
	// }()
	// go func() {
	// 	<-ch1
	// 	fmt.Println("2")
	// 	ch2 <- 2
	// }()
	// go func() {
	// 	<-ch2
	// 	fmt.Println("3")
	// 	ch3 <- 3
	// }()
	// <-ch3
}
