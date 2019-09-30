package main

import (
	"net"
	"fmt"
	"strings"
)

func main() {
	// 指定服务器 通信协议，ip地址，port端口号。 创建一个用于监听的 socket
	listener, err := net.Listen("tcp","127.0.0.1:8000")

	if err != nil {
		fmt.Println("net.listen err:",err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")
	// 阻塞监听客户端连接请求,成功建立连接，返回用于通信的 socket
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("listener.Accept err:",err)
		return
	}
	defer conn.Close()
	fmt.Println("服务器与客户端成功建立连接！！！")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:",err)
		return
	}

	readStr := string(buf[:n])
	// 处理数据-- 打印
	fmt.Println("服务器读到数据：", readStr)

	switch strings.ToLower(readStr) {

	case "are you ready?":
		conn.Write([]byte("I am Ok, Think you bor!"))
		break
	default:
		conn.Write([]byte("What happend?"))
	}

}
