package main

import (
	"fmt"
    "strings"
    "os"
)

// 环境变量

// 环境变量 是一个为 Unix 程序传递配置信息的普遍方式


var p = fmt.Println

func main() {
	
	// os.Setenv 来设置一个键值对。使用 os.Getenv 获取一个键对应的值。如果键不存在，将会返回一个空字符串
	os.Setenv("FOO", "1")
	p("FOO:", os.Getenv("FOO"))
	p("BAR:", os.Getenv("BAR"))

	p()

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		p(pair[0])
	}

}


/*
$ go build test.go

$ ./test -word=opt

$ ./test -word=opt a1 a2 a3

$ ./test -word=opt a1 a2 a3 -numb=7

$ ./test -h

$ ./test -wat

*/
