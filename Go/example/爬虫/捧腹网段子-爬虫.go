package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"strings"
	"os"
)

func httpGet3(url string) (result string, err error) {
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

func spiderContent(url string) (title,content string,err error) {
	result,err1 := httpGet3(url)
	if err1 != nil {
		err = err1
		return
	}

	// 提取title
	reg := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	titleArr := reg.FindAllStringSubmatch(result,1)

	// 提取content
	reg = regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a`)
	contentArr := reg.FindAllStringSubmatch(result,1)

	if len(titleArr)==0 {
		title = ""
	}
	if len(contentArr) == 0 {
		content = ""
	}
	title = strings.Replace(title,"&nbsp;","",-1)
	title = strings.Replace(title,"<br />","",-1)
	content = strings.Replace(content,"&nbsp;","",-1)
	content = strings.Replace(content,"<br />","",-1)

	return strings.TrimSpace(titleArr[0][1]),strings.TrimSpace(contentArr[0][1]),err
}

func SaveFile(idx int, fileMap map[string]string)  {
	//fmt.Printf("正在保存第 %d 页\n",idx)

	f, err := os.Create("第 "+ strconv.Itoa(idx) +" 页.txt")
	if err != nil {
		fmt.Println("os.Create err",err)
		return
	}
	defer f.Close()

	for title, content := range fileMap {
		f.WriteString(title + "\n\t")
		f.WriteString(content + "\n\n")
		f.WriteString("--------------------------------------\n")
	}

}

func spiderDuanZi(i int,page chan<- int) {
	fmt.Printf("正在爬取第 %d 页\n",i)

	url := "https://www.pengfu.com/xiaohua_"+ strconv.Itoa(i) +".html"

	result,err := httpGet3(url)
	if err != nil {
		fmt.Println("http get err:",err)
		return
	}

	// 提取每页 段子url
	reg := regexp.MustCompile(`<h1 class="dp-b"><a href="(.*?)"`)
	alls := reg.FindAllStringSubmatch(result,-1)

	// 创建存储 title:content 的Map
	fileMap := make(map[string]string)

	for _,link := range alls {
		title, content, err := spiderContent(link[1])
		if err != nil {
			fmt.Println("spiderContent err:",err)
			return
		}

		fileMap[title] = content
		//fmt.Println("title:",title)
		//fmt.Println("content:",content)
		//fmt.Println("")
	}
	if len(fileMap) > 0 {
		SaveFile(i,fileMap)
	}

	// 防止主go程结束
	page <- i
}

func doWorking(start, end int) {
	fmt.Printf("开始爬取第 %d 页到第 %d 页\n",start,end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go spiderDuanZi(i,page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 页 爬取完毕\n",<-page)
	}
}

func main() {
	var start,end int
	fmt.Println("请输入爬取起始页 (>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入爬取结束页 (>=start):")
	fmt.Scan(&end)

	doWorking(start,end)
}
