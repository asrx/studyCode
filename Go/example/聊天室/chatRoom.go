package main

import (
	"net"
	"fmt"
	"strings"
	"time"
)

// 创建用户结构体
type Client struct {
	C chan string
	Name string
	Addr string
}

// 创建全局map，存储在线用户
var onlineMap map[string]Client

// 创建全局 channel 传递用户消息
var message = make(chan string)

func Manager() {
	// 初始化 online
	onlineMap = make(map[string]Client)

	for{
		// 监听全局channel中是否有数据,有数据存储至 msg,无数据阻塞
		msg := <-message

		// 循环发送消息给所有在线用户
		for _,user := range onlineMap {
			user.C <- msg
		}
	}

}

func WriteMsgToClient(user Client, conn net.Conn) {
	// 监听 用户自带channel中是否有消息
	for msg := range user.C{
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMessage(user Client,msg string) string{
	buf := "【"+user.Addr+"】" + user.Name + " => " + msg

	return buf
}

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 创建一个 channel 判断用户是否活跃
	hasData := make(chan bool)

	// 创建用户结构体
	netAddr := conn.RemoteAddr().String()
	fmt.Println(netAddr," 客户端连接成功！")
	clientUser := Client{make(chan string),netAddr,netAddr}

	// 将新连接用户，添加到在线用户map中 key:IP+Port, value:client
	onlineMap[netAddr] = clientUser

	// 创建专门用来给当前 用户发送消息的go程
	go WriteMsgToClient(clientUser,conn)

	// 发送用户上线消息到 全局channel中
	//message <- "【"+netAddr+"】" + clientUser.Name + " => Login"

	message <- MakeMessage(clientUser,"Login")

	// 创建一个channel，判断用户退出状态
	isQuit := make(chan bool)

	// 创建一个匿名 go 程， 专门处理用户发送的消息
	go func() {

		for {
			buf:=make([]byte,4096)
			n,err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端:%s 已退出", clientUser.Name)
				//delete(onlineMap,clientUser.Addr)
				return
			}
			if err != nil {
				fmt.Println("conn.Read err:",err)
				return
			}

			// 读取用户输入信息
			msg := string(buf[:n-1])

			if strings.TrimSpace(msg) == "" {
				continue
			}

			// 在线用户列表
			if msg == "who" && len(msg)==3 {
				conn.Write([]byte("online User list:\n"))
				// 遍历当前Map，获取在线用户
				for _,user := range onlineMap{
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
				continue
			}

			// 改名
			if strings.HasPrefix(msg, "rename") {
				str := strings.Split(msg,"|")
				if len(str) != 2{
					conn.Write([]byte("修改昵称格式错误！\n"))
					continue
				}
				if strings.TrimSpace(str[1]) == "" {
					conn.Write([]byte("未填写昵称\n"))
					continue
				}
				oldName := clientUser.Name
				clientUser.Name = str[1]
				onlineMap[netAddr] = clientUser

				message <- MakeMessage(clientUser,oldName+",昵称修改为:"+clientUser.Name)
				conn.Write([]byte("昵称修改成功！"))
				continue
			}

			// 将读到的用户消息，写入到 message 中
			message <- MakeMessage(clientUser, msg)

			hasData <- true
		}

	}()

	// 保证不退出
	for {
		// 监听 channel 上的数据流动
		select{
		case <-isQuit:
			//delete(onlineMap,clientUser.Addr)
			//message <- MakeMessage(clientUser,"logout")
			logout(clientUser,"logout")
			return
		case <-hasData:
			// 目的是重置下面case的计时器
			continue
		case <-time.After(time.Second * 10):
			//delete(onlineMap,clientUser.Addr)
			//message <- MakeMessage(clientUser,"timeout logout")
			logout(clientUser,"timeout logout")
			return
		}
	}
}

func logout(user Client,msg string)  {
	delete(onlineMap,user.Addr)
	message <- MakeMessage(user,msg)
}

func main() {

	// 创建监听socket
	listener,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err", err)
		return
	}
	defer listener.Close()

	// 创建管理者go程，管理map和全局channel
	go Manager()

	fmt.Println("服务器端监听客户端连接...")
	for {	// 循环监听客户端请求
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Accept err", err)
			return
		}

		// 启动go程，处理客户端数据请求
		go HandlerConnect(conn)
	}
}
