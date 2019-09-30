package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
)

const PATH  = "./test.txt"

func read1()  {
	buf, err := ioutil.ReadFile(PATH)
	if err != nil {
		panic(err)
	}
	//fmt.Println(buf)
	str := string(buf)
	fmt.Println(str)
}

func read2()  {
	file,err := os.Open(PATH)
	if err != nil {
		panic(err)
	}

	buf:=make([]byte,4096)
	str := ""

	for{
		n,err := file.Read(buf)
		if n == 0 {
			fmt.Println("Read Done.\n")
			break
		}
		if err != nil {
			panic(err)
		}

		str += string(buf[:n])
		//fmt.Println(str)
	}

	fmt.Println(str)
}

func read3() {
	f,_ := os.Open(PATH)

	reader := bufio.NewReader(f)
	//buf,_,_ := reader.ReadLine()
	//buf,_ := reader.Peek(8)
	buf := make([]byte,8)
	for{
		n, _ := reader.Read(buf)

		if n == 0 {
			fmt.Println("Read Done.\n")
			break
		}

		fmt.Print(string(buf[:n]))
	}

	//fmt.Println(string(buf))
}

func main() {
	read3()
}
