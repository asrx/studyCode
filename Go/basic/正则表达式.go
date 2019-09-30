package main

import (
	"fmt"
	"bytes"
	"regexp"
)

// 正则表达式
// http://golang.org/pkg/regexp/

var p = fmt.Println

func main() {

	// 测试一个字符串是否符合一个表达式
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match)

	// Compile 一个优化的 Regexp 结构体
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 匹配测试
	p(r.MatchString("peach"))	

	// 查找匹配字符串
	p(r.FindString("peach punch"))

	// 查找第一次匹配的字符串的，但是返回的匹配开 始和结束位置索引，而不是匹配的内容
	p(r.FindStringIndex("peach punch"))

	// 返回完全匹配和局部匹配的字符串。例如， 这里会返回 p([a-z]+)ch 和 `([a-z]+) 的信息
	p(r.FindStringSubmatch("peach punch"))

	// 返回完全匹配和局部匹配的索引位置
	p(r.FindStringSubmatchIndex("peach punch"))

	// 返回所有的匹配项，而不仅仅是首 次匹配项。例如查找匹配表达式的所有项
	p(r.FindAllString("peach punch pinch",-1))

	p(r.FindAllStringSubmatchIndex("peach punch pinch",-1))
	
	// 函数提供一个正整数来限制匹配次数
	p(r.FindAllString("peach punch pinch", 2))

	// 提供 []byte 参数并将 String 从函数命中去掉
	p(r.Match([]byte("peach")))

	// 创建正则表达式常量时，可以使用 Compile 的变体 MustCompile 。因为 Compile 返回两个值，不能用于常量
	r = regexp.MustCompile("p([a-z]+)ch")
	p(r)

	// regexp 包也可以用来替换部分字符串为其他值
	p(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变量允许传递匹配内容到一个给定的函数中
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out))
}


/*

*/
