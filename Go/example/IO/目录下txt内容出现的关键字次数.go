package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"io"
)

func checkErr3(err error){
	if err != nil {
		fmt.Println("Err:",err)
		panic(err)
	}
}

/**
	读取指定目录下的所有txt中love出现的次数
 */
func main() {

	// 获取用户输入的目录路径
	fmt.Println("请输入路径:")
	//var path string
	//fmt.Scan(&path)

	path := "./"

	// 获取目录下所有内容
	f, err := os.OpenFile(path,os.O_RDONLY,os.ModeDir)
	checkErr3(err)

	defer f.Close()

	ArrLove := 0

	// 获取所有文件
	info,err := f.Readdir(-1)
	checkErr3(err)

	for _,fileInfo := range info{

		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(),".txt") {
			fmt.Println(path+fileInfo.Name())

			//fr, err := os.OpenFile(path+fileInfo.Name(),os.O_RDWR,6)
			fr, err := os.Open(path+fileInfo.Name())
			checkErr3(err)
			defer fr.Close()

			//buf = make([]byte, 1024)
			reader := bufio.NewReader(fr)

			for {
				//buf, err := reader.ReadBytes('\n')
				buf, _, err := reader.ReadLine()

				if err != nil && err == io.EOF {
					fmt.Println("Read Done.")
					break
				}else if err != nil{
					panic(err)
				}

				words := strings.Fields(string(buf))

				for _,w := range words{
					if strings.ToLower(w) == "love" {
						ArrLove++
					}
				}
			}



		}
	}

	fmt.Println("love 共出现次数:", ArrLove)
}
