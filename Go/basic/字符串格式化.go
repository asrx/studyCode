package main

import (
	_"strings"
	_"sort"
	"fmt"
	"os"
)

var p = fmt.Printf

// 字符串格式化

type point struct{
	x, y int
}

func main() {

	pt := point{1, 2}
	p("%v\n", pt)

	// %+v 的格式化输出内容将包括 结构体的字段名
	p("%+v\n", pt)
	
	// %#v 形式则输出这个值的 Go 语法表示。例如，值的 运行源代码片段
	p("%#v\n", pt)

	// 需要打印值的类型，使用 %T
	p("%T\n", pt)

	// 格式化布尔类型
	p("%t\n", true)

	// 格式化整型数有多种方式，
	// 使用 %d进行标准的十进 制格式化
	p("%d\n",123)

	// 二进制
	p("%b\n",14)

	// 输出给定整数的对应字符
	p("%c\n", 33)

	// 十六进制
	p("%x\n",456)

	// 浮点型
	p("%f\n",78.9)

	// %e 和 %E 将浮点型格式化为（稍微有一点不 同的）科学记数法表示形式
	p("%e\n",123400000.0)
	p("%E\n",123400000.0)

	// %s字符串
	p("%s\n", "\"string\"")

	// 像 Go 源代码中那样带有双引号的输出，使用 %q
	p("%q\n", "\"string\"")	

	// 和上面的整型数一样，%x 输出使用 base-16 编码的字 符串，每个字节使用 2 个字符表示
	p("%x\n", "hex this")	

	//输出一个指针的值，使用 %p	
	p("%p\n", &pt)	

	// 使用在 % 后面使用数字来控制输出宽度
	p("|%6d|%6d|\n", 12, 345)

	// 指定浮点型的输出宽度，同时也可以通过 宽度. 精度 的语法来指定输出的精度
	p("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// 左对齐，使用 - 标志
	p("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// Sprintf 则格式化并返回一个字 符串而不带任何输出
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	// 使用 Fprintf 来格式化并输出到 io.Writers 而不是 os.Stdout
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}


/*

*/
