package main

import "fmt"

// 死锁1
func main_ss1() {

	ch := make(chan int)
	ch <-700
	num := <-ch
	fmt.Println("num = ",num)
}

// 死锁2
func main_ss2() {

	ch := make(chan int)
	num := <-ch
	fmt.Println("num = ",num)

	go func() {
		ch <- 789
	}()

}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for {
			select {
			case num:=<-ch1:
				ch2<-num
			}
		}
	}()

	for {
		select {
		case num:=<-ch2:
			ch1<-num
		}
	}
}