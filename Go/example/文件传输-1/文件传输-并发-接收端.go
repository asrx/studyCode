package main

import (
	"net"
	"fmt"
	"io"
	"os"
)

func receiveFile2(conn net.Conn,fileName string)  {
	// 获取连接客户端的 Addr
	addr := conn.RemoteAddr()
	fmt.Println(addr, " 客户端成功连接！")

	// 按照文件名创建新文件
	f, err := os.Create(fileName)
	if err != nil{
		fmt.Println("os.Create err:",err)
		return
	}
	defer f.Close()

	// 从网络中读取数据，写入文本文件
	buf:=make([]byte,4096)
	for{
		n, _ := conn.Read(buf)
		if n==0 {
			fmt.Println("文件接收完成")
			return
		}
		// 写入文件
		f.Write(buf[:n])

	}
}

func main() {
	// 创建监听socket
	listener,err := net.Listen("tcp","127.0.0.1:8008")
	if err != nil {
		fmt.Println("net.Listen err:",err)
		return
	}
	defer listener.Close()

	// 阻塞等待发送端连接 Accept()
	fmt.Println("阻塞等待发送端连接 Accept()...")

	for{
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:",err)
			return
		}

		//go HandlerReceive(conn)

		// 接收（读取）文件名，保存
		buf := make([]byte, 4096)

		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完成！")
			}else {
				fmt.Println("conn.Read err:",err)
			}
			return
		}

		fileName := string(buf[:n])
		fmt.Println("预接收文件名：", fileName)

		// 回发 "ok" 给 发送端
		conn.Write([]byte("ok"))


		// 接收文件内容，写入文件中
		go receiveFile2(conn,fileName)
	}

}

