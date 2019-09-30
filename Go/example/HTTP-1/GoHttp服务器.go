package main

import (
	"fmt"
	"os"
	"net/http"
)

func errFunc3(err error,info string) {
	if err != nil {
		fmt.Println(info,err)
		os.Exit(1)	// 将当前进程结束
	}
}


func myHandler(writer http.ResponseWriter, request *http.Request) {
	// writer 写会给客户端（浏览器）的数据

	// request 从客户端（浏览器）读到的数据

	writer.Write([]byte("this is a Web Server"))

	fmt.Println("Header:",request.Header)
	fmt.Println("URL:",request.URL)
	fmt.Println("Method:",request.Method)
	fmt.Println("Host:",request.Host)
	fmt.Println("RemoteAddr:",request.RemoteAddr)
	fmt.Println("Body:",request.Body)

}

func main() {
	// 注册回调函数，该函数在客户端访问服务器时会，自动被调用
	http.HandleFunc("/",myHandler)

	http.ListenAndServe("127.0.0.1:8000",nil)
}
