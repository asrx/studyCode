package main

import (
	"fmt"
	"strconv"
)

// 数字解析

var p = fmt.Println

func main() {

	// ParseFloat 解析浮点数，这里的 64 表示解 析的数的位数
	f, _ := strconv.ParseFloat("1.234", 64)
	p(f)

	// ParseInt 解析整型数时，例子中的参数 0 表 示自动推断字符串所表示的数字的进制。64 表示返回的 整型数是以 64 位存储的
	i, _ := strconv.ParseInt("123", 0, 64)
	p(i)


	// ParseInt 会自动识别出十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	p(d)

	// ParseUint 也是可用的
	u, _ := strconv.ParseUint("789", 0, 64)
	p(u)

	// Atoi 是一个基础的 10 进制整型数转换函数
	k, _ := strconv.Atoi("135")
	p(k)

	// 在输入错误时，解析函数会返回一个错误
	_, e := strconv.Atoi("wat")
	p(e)
}


/*

*/
