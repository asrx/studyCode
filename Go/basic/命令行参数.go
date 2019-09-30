package main

import (
	"fmt"
    "os"
)

// 命令行参数

var p = fmt.Println

func main() {
	
	// os.Args 提供原始命令行参数访问功能。
	// 注意，切片中 的第一个参数是该程序的路径，
	// 且 os.Args[1:]保存 所有程序的参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 可以使用标准的索引位置方式取得单个参数的值
	arg := os.Args[3]

	p(argsWithProg)
	p(argsWithoutProg)
	p(arg)
}

/*
$ go build test.go
$ ./test a b c d
*/
