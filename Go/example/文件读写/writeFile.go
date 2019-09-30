package main

import (
	"io/ioutil"
	"os"
	"bufio"
	"fmt"
)

const PATH2  = "./test2.txt"

func write1() {
	err := ioutil.WriteFile(PATH2,[]byte("Test write file\n\nI am Jerry.Xu"),0644)
	if err != nil {
		panic(err)
	}
}

func write2() {
	file,_ := os.Create(PATH2)
	defer file.Close()

	file.Write([]byte("Test 2 write file\n\nI am Jerry.Xu\n"))

	file.WriteString("哇哈哈哈")

	file.Sync()
}

func write3()  {
	file,_:=os.Create(PATH2)
	defer file.Close()

	writer := bufio.NewWriter(file)
	n,_ := writer.WriteString("Golang 的文件读取方法很多,刚上手时不知道怎么选择,所以贴在此处便后速查。 一次性读取 小文件推荐一次性读取,这样程序更简单,而且速度最快。 代码如...")
	writer.Flush()
	
	fmt.Println(n)
}

func main() {
	//write1()

	//write2()

	write3()
}
