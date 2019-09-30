package main

import (
	"fmt"
	"bufio"
    "io/ioutil"
    "os"
)

// 写文件

var p = fmt.Println


// 精简下面 的错误检查过程
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")

	path := "/Users/Jerry/Desktop/Test/go/dat1"

	// 开始，这里是展示如何写入一个字符串（或者只是一些 字节）到一个文件
	err := ioutil.WriteFile(path,d1,0644)
	check(err)

	// 更细粒度的写入，先打开一个文件
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	// 写入你想写入的字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)


	// WriteString
	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 Sync 来将缓冲区的信息写入磁盘
	f.Sync()


	// bufio 提供了和我们前面看到的带缓冲的读取器一 样的带缓冲的写入器
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}

/*

*/
