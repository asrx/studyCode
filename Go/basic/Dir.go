package main

import (
	"fmt"
	"os"
	"strings"
)


func checkErr2(err error){
	if err != nil {
		fmt.Println("Err:",err)
		panic(err)
	}
}

func main() {

	// 获取用户输入的目录路径
	fmt.Println("请输入待查询的目录:")
	var path string
	fmt.Scan(&path)

	f, err := os.OpenFile(path,os.O_RDONLY, os.ModeDir)
	checkErr2(err)

	defer f.Close()

	// 读取目录项
	info, err := f.Readdir(-1)	// -1 读取目录所有项
	checkErr(err)

	for _, fileInfo := range info{

		//if fileInfo.IsDir() {
		//	fmt.Println(fileInfo.Name()," 是一个目录")
		//}else {
		//	fmt.Println(fileInfo.Name()," 是一个文件")
		//}

		if !fileInfo.IsDir() {

			if strings.HasSuffix(fileInfo.Name(), ".html") {
				//fmt.Println("txt 文件有:")
				//fmt.Println("html 文件有:")
				fmt.Println(fileInfo.Name())
			}
		}
	}

}
