package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func checkErr(err error){
	if err != nil {
		fmt.Println("Err:",err)
		panic(err)
	}
}

const fileName = "./test.txt"

func main_CreateFile() {



	file, err := os.Create(fileName)
	checkErr(err)

	defer file.Close()

	fmt.Println("successful")
}

func main_OpenFIle() {

	file, err := os.Open(fileName)

	checkErr(err)

	defer file.Close()

	_, err = file.WriteString("write...")
	checkErr(err)

	fmt.Println("successful")
}

func main_Write() {

	file, err := os.OpenFile(fileName,os.O_RDWR, 6)

	checkErr(err)

	defer file.Close()

	_, err = file.WriteString("helloworld\r\n")
	checkErr(err)

	n,_  := file.Seek(-5,io.SeekEnd)

	//file.WriteString("xxx")
	file.WriteAt([]byte("1111"), n)

	fmt.Println("n:", n)

	fmt.Println("successful")
}

func main() {

	file, err := os.OpenFile(fileName,os.O_RDWR, 6)

	checkErr(err)

	defer file.Close()

	// 创建一个带有缓冲区的Reader
	reader := bufio.NewReader(file)

	checkErr(err)
	defer file.Close()

	for {
		buf, err := reader.ReadBytes('\n')	// 读一行数据
		if err != nil && err == io.EOF {
			fmt.Println("文件读写完成！")
			return
		}else if err != nil {
			panic(err)
		}

		fmt.Print(string(buf))

	}


}