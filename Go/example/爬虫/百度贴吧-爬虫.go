package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"os"
)

func httpGet1(url string) (result string,err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte,4096)
	for {
		n,err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网页完成！\n")
			break
		}
		if err2 != nil && err2 != io.EOF{
			err = err2
			return
		}

		// 累加每次循环读到的 buf 数据，存入 result 一次性返还
		result += string(buf[:n])
	}
	return
}

func working1(start, end int) {
	fmt.Printf("正在爬取第 %d 页到 %d 页.....\n", start,end)

	// 循环爬取每一页
	for i := start; i <= end; i++ {
		fmt.Println("正在爬取第 "+ strconv.Itoa(i) +" 页")

		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1) * 50)
		result, err:= httpGet1(url)
		if err != nil {
			fmt.Println("httpGet err:",err)
			continue
		}



		// 将读到的整网页数据，保存一个文件
		f, err := os.Create("第 "+strconv.Itoa(i)+" 页.html")
		if err != nil {
			panic(err)
		}

		f.WriteString(result)
		f.Close()			// 保存好一个文件，关闭一个文件
	}

}
func main() {

	/*
	https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
	https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=50
	https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=100
	https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=150
	*/
	// 指定爬取起始页、终止页
	var start,end int
	fmt.Println("输入爬取的起始页（>=1）")
	fmt.Scan(&start)
	fmt.Println("输入爬取的终止页（>=start）")
	fmt.Scan(&end)

	working1(start,end)
}
