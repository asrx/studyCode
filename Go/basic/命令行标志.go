package main

import (
	"fmt"
    "flag"
)

// 命令行标志

var p = fmt.Println

func main() {
	
	// 声明一个默认值为 "foo" 的字符串标志 word 并带有一个简短的描述
	// flag.String 函数返回一个字 符串指针（不是一个字符串值）
	wordPtr := flag.String("word","foo","a string")

	// 声明 numb 和 fork 标志
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	// 用程序中已有的参数来声明一个标志也是可以的。注 意在标志声明函数中需要使用该参数的指针
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")


	// 所有标志都声明完成以后，调用 flag.Parse() 来执行 命令行解析
	flag.Parse()

	p("word:", *wordPtr)
	p("numb:", *numbPtr)
	p("fork:", *boolPtr)
	p("svar:", svar)
	p("tail:", flag.Args())

}


/*
$ go build test.go

$ ./test -word=opt

$ ./test -word=opt a1 a2 a3

$ ./test -word=opt a1 a2 a3 -numb=7

$ ./test -h

$ ./test -wat

*/
