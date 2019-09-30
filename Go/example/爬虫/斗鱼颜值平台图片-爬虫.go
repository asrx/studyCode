package main

import (
	"net/http"
	"io"
	"fmt"
	"regexp"
	"os"
	"strconv"
)

func httpGet4(url string) (result string, err error) {
	resp,err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for{
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
		}
		result += string(buf[:n])
	}

	return
}

func saveImage(idx int, imageUrl string,finsh chan<- int) {
	fmt.Printf("正在保存第 %d 张图片\n",idx)

	// 读取网络图片
	resp, err := http.Get(imageUrl)
	if err != nil {
		fmt.Println("http.Get err:",err)
		return
	}
	defer resp.Body.Close()

	// 创建图片
	f,err := os.Create(strconv.Itoa(idx) + ".jpg")
	if err != nil {
		fmt.Printf("os.Create err:",err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)

	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("resp.Body.Read err:",err)
			return
		}

		f.Write(buf[:n])
	}

	finsh <- idx
}

func main() {
	url := "https://www.douyu.com/g_yz"

	finsh := make(chan int)

	result,err := httpGet4(url)
	if err != nil {
		fmt.Println("http get err:",err)
	}

	// 图片
	reg := regexp.MustCompile(`"rs16":"(?s:(.*?))"`)
	alls := reg.FindAllStringSubmatch(result,-1)

	for idx,img := range alls{
		go saveImage((idx+1),img[1],finsh)
	}

	// 防止主go程退出
	for i := 0; i < len(alls);i++{
		fmt.Printf("第%d张照片保存完成\n",<-finsh)
	}
}
