package main

import "fmt"


// Go 语言指针作为函数参数

func main() {

	var a int = 100
	var b int = 200

	fmt.Printf("交换前 a 的值 : %d\n", a )
	fmt.Printf("交换前 b 的值 : %d\n", b )

	swap(&a,&b)

	fmt.Printf("交换后 a 的值 : %d\n", a )
	fmt.Printf("交换后 b 的值 : %d\n", b )

	
}

func swap(x *int, y *int){
	var temp int

	temp = *x
	*x = *y
	*y = temp
}
