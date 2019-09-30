package main

import (
	"os"
	"io"
	"fmt"
	"math/rand"
	"time"
	"io/ioutil"
)

func RandomSource(count int) <-chan int {
	out := make(chan int,1024)

	go func() {
		for i := 0; i < count; i++ {
			out<-rand.Int()
		}
		close(out)
	}()
	return out
}

var startTime time.Time

func Init() {
	startTime = time.Now()
}

func cp1(originalPath,targetPath string) {
	buf := make([]byte,4096)
	f,err := os.Open(originalPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	writer,err := os.Create(targetPath)
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	for{
		n,err := f.Read(buf)
		if n == 0 {
			fmt.Println("读取完成！")
			return
		}
		if err != nil && err != io.EOF {
			panic(err)
		}
		writer.Write(buf[:n])
	}
	fmt.Println("cp1 Done.")
}

func cp2(originalPath,targetPath string)  {
	buf,err := ioutil.ReadFile(originalPath)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(targetPath,buf,0644)
	if err != nil{
		panic(err)
	}
	fmt.Println("cp2 Done.")
}

func cp3(originalPath,targetPath string)  {

	//io.Copy(targetPath,originalPath)
	fileStat,err := os.Stat(originalPath)
	if err != nil {
		panic(err)
	}

	if !fileStat.Mode().IsRegular() {
		fmt.Printf("%s,不是常规文件",originalPath)
		return
	}

	originalFile,err := os.Open(originalPath)
	if err != nil {
		panic(err)
	}
	defer originalFile.Close()

	targetFile,err := os.Create(targetPath)
	if err != nil {
		panic(err)
	}
	defer targetFile.Close()

	n,err:= io.Copy(targetFile,originalFile)

	fmt.Println("文件大小: ",n)
	fmt.Println("cp3 Done.")
}

func main() {
	// 文件原始路径
	originalPath := "/Users/Jerry/Desktop/test.txt"
	// 拷贝路径
	targetPath := "./test.txt"
	Init()
	
	//cp1(originalPath,targetPath)

	cp2(originalPath,targetPath)

	//cp3(originalPath,targetPath)

	fmt.Println(time.Now().Sub(startTime))
}
