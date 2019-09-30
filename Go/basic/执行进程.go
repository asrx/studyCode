package main

import (
	"syscall"
    "os"
    "os/exec"
)

// 执行进程

/*

*/

func main() {
	
	// 将执行 ls 命令。Go 需要提供我 们需要执行的可执行文件的绝对路径，所以我们将使用 exec.LookPath 来得到它（大概是 /bin/ls）
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec 需要的参数是切片的形式的（不是放在一起的一个大字 符串）。我们给 ls 一些基本的参数。注意，第一个参数需要 是程序名
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec 同样需要使用环境变量。 这里我们仅提供当前的环境变量
	env := os.Environ()

	// 这里是 syscall.Exec 调用。如果这个调用成功，那么我们的 进程将在这里被替换成 /bin/ls -a -l -h 进程。
	// 如果存 在错误，那么我们将会得到一个返回值
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}


/*
注意 Go 并不提供一个经典的 Unix fork 函数。通常这不 是个问题，因为运行 Go 协程，生成进程和执行进程覆盖了 fork 的大多数使用场景

*/
