package main

import "fmt"

func main() {

	fmt.Println("********【标准】********")

	// 标准
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	fmt.Println("********【一分支多值】********")

	// 一分支多值
	a = "num"
	switch a {
	case "num","daddy":
		fmt.Println("family")
	}


	fmt.Println("********【分支表达式】********")

	// 分支表达式

	var r int =11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)
	}

	fmt.Println("********【跨越case的fallthrough】********")

	var s = "hello"
	switch {
	case s == "hello":
		fmt.Println("hello")
		fallthrough
	case s != "world":
		fmt.Println("world")
	}

	//fmt.Println(reflect.TypeOf(l))
}
