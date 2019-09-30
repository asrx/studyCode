package main

import "fmt"

// 切片


func main() {
	
	// var numbers = make([]int, 3, 5)

	// printSlice(numbers)

	numbers := []int{0,1,2,3,4,5,6,7,8}
	printSlice(numbers)

	fmt.Println("numbers ==", numbers)

	fmt.Println("numbers[1:4] ==",numbers[1:4])

	//
	numbers = append(numbers,1)

	numbers = append(numbers,2,3,4)

	printSlice(numbers)

	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	printSlice(numbers1)	

}


func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x),cap(x),x)
	
}
