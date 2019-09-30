package main

import "net/http"

func handler(writer http.ResponseWriter, request *http.Request) {
	// writer 写会给客户端（浏览器）的数据

	// request 从客户端（浏览器）读到的数据

	writer.Write([]byte("Hello Wrold"))
}

func main() {
	// 注册回调函数，该回调函数会在服务器被访问时，自动被调用
	http.HandleFunc("/itcast",
		handler)

	// 绑定服务器地址
	http.ListenAndServe("127.0.0.1:8000",nil)
}
