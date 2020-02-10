package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

func checkErr(e error){
	if e != nil {
		panic(e)
	}
}
func main() {

	// 将文本内容读取到内存中
	dat,err:= ioutil.ReadFile("./dat")
	checkErr(err)
	fmt.Println(string(dat))


	f, err:= os.Open("./dat")
	checkErr(err)
	defer f.Close()

	// 从文件开始位置读取一节字节，这里最多读取5个字节
	b1:=make([]byte,5)
	n1,err:=f.Read(b1)
	checkErr(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// `Seek` 到一个文件已知的位置并从这个位置开始进行读取
	o2,err:=f.Seek(6,0)
	checkErr(err)
	b2:=make([]byte,2)
	n2,err:=f.Read(b2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n",n2,o2,string(b2))

	// ReadAtLeast
	o3,err:=f.Seek(6,0)
	checkErr(err)
	b3:=make([]byte,2)
	n3,err:=io.ReadAtLeast(f,b3,2)
	checkErr(err)
	fmt.Printf("%d bytes @ %d: %s\n",n3,o3,string(b3))

	_, err = f.Seek(0,0)
	checkErr(err)

	// `bufio` 包实现了带缓冲的读取，提供读取性能和很多附加读取函数
	r4 := bufio.NewReader(f)
	b4,err:=r4.Peek(5)
	checkErr(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}
