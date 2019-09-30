package main

import (
	"net"
	"fmt"
	"time"
)

func main() {

	// 组织一个 udp 地址结构, 指定服务器的IP+port
	serverAddr,err  := net.ResolveUDPAddr("udp","127.0.0.1:8003")
	if err != nil {
		fmt.Println("net.ResolveUDPAddr err:",err)
		return
	}
	fmt.Println("udp 服务器地址结构，创建完成!!!")

	// 创建用于通信的 socket
	udpConn, err := net.ListenUDP("udp",serverAddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:",err)
		return
	}
	defer udpConn.Close()
	fmt.Println("udp 服务器通信 socket 创建完成!!!")

	// 读取客户端发送的数据
	buf := make([]byte, 4096)

	for {
		// 返回: 读取的字节数，客户端地址，error
		n,clientAddr,err := udpConn.ReadFromUDP(buf)	// 主go程读取客户端发送数据
		if err != nil {
			fmt.Println("udpConn.ReadFrom err:",err)
			return
		}

		// 模拟处理数据
		fmt.Printf("服务器读到 %v 的数据：%s\n",clientAddr, string(buf[:n]))

		go func() {					// 每有一个客户端连接上来，启动一个go程，写数据
			// 提取系统当前时间
			dayTime := time.Now().String()

			// 回写数据给客户端
			_, err = udpConn.WriteToUDP([]byte(dayTime+"\n"), clientAddr)
			if err != nil {
				fmt.Println("udpConn.WriteToUDP err:",err)
				return
			}
		}()
	}
}
