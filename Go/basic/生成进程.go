package main

import (
	"fmt"
    "io/ioutil"
    "os/exec"
)

// 生成进程

/*
有时，我们的 Go 程序需要生成其他的，非 Go 进程。例如，这个 网站的语法高亮是通过在 Go 程序中生成一个 pygmentize 来实现的
*/

var p = fmt.Println

func main() {
	
	// exec.Command 函数帮助我们创 建一个表示这个外部进程的对象
	dateCmd := exec.Command("date")

	// .Output 是另一个帮助我们处理运行一个命令的常见情况 的函数，它等待命令运行完成，并收集命令的输出。如果没 有出错，dateOut 将获取到日期信息的字节
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	p("> date")
	p(string(dateOut))


	// 从外部进程的 stdin 输入数据并从 stdout 收集结果
	grepCmd := exec.Command("grep", "hello")

	// 明确的获取输入/输出管道，运行这个进程，写入 一些输入信息，读取输出的结果，最后等待程序运行结束
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	p("> grep hello")
	p(string(grepBytes))

	// 在生成命令时，我们需要提供显式描述的命令和参数 数组，而不能只传递一个命令行字符串。如果你想使用一个 字符串生成一个完整的命令，那么你可以使用 bash 命令 的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err !=nil {
		panic(err)
	}
	p("> ls -a -l -h")
	p(string(lsOut))


}


/*

*/
