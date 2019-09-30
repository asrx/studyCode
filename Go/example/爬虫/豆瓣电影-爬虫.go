package main

import (
	"fmt"
	"strconv"
	"net/http"
	"io"
	"regexp"
	"os"
)

func http_Get(url string) (result string,err error){
	resp, err1 := http.Get(url)
	if err1 != nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)

		if n == 0 {
			break
		}

		if err2 != nil && err2 != io.EOF{
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return
}

func SaveToFile(filmName,filmScore,peopleNum [][]string,idx int){

	f,err := os.Create("第 " + strconv.Itoa(idx) + " 页.txt")
	if err != nil {
		fmt.Println("os.Create err:",err)
		return
	}

	f.WriteString("电影名称\t\t\t\t评分\t\t\t人数\n")
	for i := 0; i< len(filmName); i++  {
		f.WriteString(filmName[i][1] +
			"\t\t\t" +
			filmScore[i][1] +"\t\t"+
			peopleNum[i][1] + "\n")
	}

	f.Close()
}

// 爬取一个豆瓣页面信息
func SpiderDoubanMovie(i int, page chan<- int) {
	fmt.Println("正在爬取第 "+ strconv.Itoa(i) +" 页")

	url := "https://movie.douban.com/top250?start="+ strconv.Itoa((i-1) * 25) +"&filter="

	result,err := http_Get(url)
	if err != nil {
		fmt.Println("http get err:",err)
		return
	}

	//fmt.Println(result)

	// 解析、编译正则表达式
	pattern1:= `<img width="100" alt="(?s:(.*?))"`
	rep1 := regexp.MustCompile(pattern1)

	// 提取信息
	filmName := rep1.FindAllStringSubmatch(result,-1)

	// 分数
	pattern2 := `<span class="rating_num" property="v:average">(?s:(.*?))</span>`
	rep2 := regexp.MustCompile(pattern2)
	filmScore := rep2.FindAllStringSubmatch(result,-1)

	// 评论人数
	//pattern3 := `<span>(?s:(\d*?))人评价</span>`
	pattern3 := `<span>(\d*?)人评价</span>`
	rep3 := regexp.MustCompile(pattern3)
	peopleNum := rep3.FindAllStringSubmatch(result,-1)

	SaveToFile(filmName,filmScore,peopleNum,i)

	page <- i
}

func toWork(start, end int) {
	fmt.Printf("正在爬取第 %d 页到 %d 页.....\n", start,end)

	page := make(chan int)

	for i := start; i <= end; i++ {
		go SpiderDoubanMovie(i,page)
	}

	for i := start; i <= end; i++ {
		fmt.Printf("第 %d 个页面爬取完成\n",<-page)
	}


}

func main() {
	// 横向
	/*
	https://movie.douban.com/top250?start=0&filter=
	https://movie.douban.com/top250?start=25&filter=
	https://movie.douban.com/top250?start=50&filter=
	https://movie.douban.com/top250?start=75&filter=
	*/

	// 纵向
	/*
	电影名称:`<img width="100" alt="(?s:(.*?))"`
	评分:`<span class="rating_num" property="v:average">(?s:(.*?))</span>`
	评分人数:`<span>(?s:(.*?))人评价</span>`
	*/


	var start,end int
	fmt.Println("请输入爬取的起始页 (>=1):")
	fmt.Scan(&start)


	fmt.Println("输入爬取的终止页（>=start）")
	fmt.Scan(&end)

	toWork(start,end)

}
