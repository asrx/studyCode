package main

import (
	"fmt"
	"regexp"
)

func main() {
	//str := "abc a7c mfc cat aMc 8ca asc cba"

	str := "3.14 123.123 .68 haha 1.0 abc 7. ab.3 66.6 123."

	// 解析、编译正则表达式
	//rp := regexp.MustCompile(`a.c`)	//`` 表示使用原生字符串

	//rp := regexp.MustCompile(`a[^0-9a-z]c`)

	//rp := regexp.MustCompile(`[0-9]+\.[0-9]+`)
	rp := regexp.MustCompile(`\d+\.\d+`)

	// 提取信息
	alls := rp.FindAllStringSubmatch(str,-1)


	fmt.Println("alls:",alls)


}
