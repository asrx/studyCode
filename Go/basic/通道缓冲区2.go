package main

import (
	"fmt"
	_"errors"
	_"strconv"
	_"math"
	_"time"
)

func main() {

	message := make(chan string)

	// go func () {
	// 	message <- "ping"
	// }()

	msg := <-message
	fmt.Println(msg)


	//
	message2 := make(chan string, 2)

	go func () {
		message2 <- "buffered"	
	}()

	go func () {
		message2 <- "channel"	
	}()
	

	fmt.Println(<-message2)
	fmt.Println(<-message2)

}
