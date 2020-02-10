package main

import (
	"os/exec"
	"os"
	"syscall"
)

func main() {
	binary, _ := exec.LookPath("ls")

	args := []string{"ls","-a","-l","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary,args,env)
	if execErr!=nil{
		panic(execErr)
	}
}
