package main

import (
	"regexp"
	"fmt"
)

func main() {
	str := `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>
<title>Go语言</title>
<div>hello regexp</div>
<div>hello 222</div>
<div>hello 888</div>
<div>hello 666</div>
<div>
	青春献给小酒桌
	醉生梦死就是喝
	一个人没有很好地做过孩子，他就无法很好地做一个成人。—— 卢悦
</div>
<body>身体</body>

<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
</frameset>
</html>`

	//rep := regexp.MustCompile(`<div>(.*)</div>`)
	//模式修饰符?s: 【.】可以匹配\n
	rep := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	alls := rep.FindAllStringSubmatch(str,-1)

	for _, one := range alls{
		fmt.Println("one[0]：",one[0])
		fmt.Println("one[1]：",one[1])
	}
	//fmt.Println("alls：",alls)
}
