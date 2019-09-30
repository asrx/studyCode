package main

import (
	"net"
	"fmt"
	"strings"
	"io"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 获取连接客户端的 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, " 客户端成功连接！")

	// 循环读取客户端数据
	buf := make([]byte, 4096)
	for{
		n, err := conn.Read(buf)

		//fmt.Println(buf[:n])
		if "exit\n" == string(buf[:n]) || "exit\r\n" == string(buf[:n]) {
			fmt.Println("服务器接收到客户端退出请求，服务器关闭！")
			return
		}

		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭，服务器断开连接！！！")
			return
		}

		if err != nil && err == io.EOF{
			fmt.Println("conn.Read Will Close")
			return
		}else if err != nil {
			fmt.Println("conn.Read err:",err)
			return
		}

		str := string(buf[:n])
		fmt.Println("服务器读到的数据", str)

		// 回发客户端（小写转大写）
		conn.Write([]byte(strings.ToUpper(str)))
	}
}
func main() {

	// 创建监听套接字
	listener, err := net.Listen("tcp","127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.listen err:",err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器等待客户端建立连接...")

	// 监听客户端连接请求
	for{
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:",err)
			return
		}

		// 具体完成服务器和客户端的数据通信
		go HandlerConnect(conn)
	}


}
