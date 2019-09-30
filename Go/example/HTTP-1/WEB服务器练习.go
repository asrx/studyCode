package main

import (
	"net/http"
	"fmt"
	"os"
)

func OpenSendFile(fName string,w http.ResponseWriter) {
	pathFileName := "/Users/Jerry/Pictures/"+fName

	f,err := os.Open(pathFileName)

	if err != nil {
		fmt.Println("Open err:",err)
		w.Write([]byte("No such file or directory !"))
		return
	}
	defer f.Close()

	buf := make([]byte,4096)
	for{
		n,_:= f.Read(buf)
		if n == 0 {
			return
		}
		w.Write(buf[:n])
	}

}

func myHandler2(writer http.ResponseWriter, request *http.Request) {
	// writer 写会给客户端（浏览器）的数据

	// request 从客户端（浏览器）读到的数据

	writer.Write([]byte("Hello Wrold"))

	fmt.Println("客户端请求：",request.URL)

	OpenSendFile(request.URL.String(),writer)
}

func main() {

	http.HandleFunc("/", myHandler2)

	http.ListenAndServe("127.0.0.1:8000",nil)
}
