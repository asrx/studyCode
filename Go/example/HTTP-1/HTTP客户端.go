package main

import (
	"fmt"
	"net/http"
	"io"
)

func main() {

	resp,err := http.Get("https://www.baidu.com")
	//resp,err := http.Get("http://www.itcast.cn")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 查看应答包
	fmt.Println("Header:",resp.Header)
	fmt.Println("Status:",resp.Status)
	fmt.Println("StatusCode:",resp.StatusCode)
	fmt.Println("Proto:",resp.Proto)

	buf := make([]byte,4096)
	var result string
	for{
		n,_ := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("--------Read finsh!")
			break
		}
		if err != nil && err != io.EOF{
			fmt.Println("Body.Read err:",err)
			break
		}
		result += string(buf[:n])
	}

	fmt.Println(result)
}
