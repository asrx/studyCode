package main

import (
	"os"
	"bufio"
	"../../pipeline"
	"fmt"
	"strconv"
)

func main() {
	//fileName := "small.in"
	//outFileName := "small.out"
	//p := createPipeline(fileName,512,4)

	//fileName := "large.in"
	//outFileName := "large.out"
	//p := createPipeline(fileName,800000000,4)


	// Network
	fileName := "large.in"
	outFileName := "large.out"
	p := createNetworkPipeline(fileName,800000000,4)

	writeToFile(p, outFileName)
	printFile(outFileName)
}

func printFile(filename string)  {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file,-1)
	count := 0
	for v:=range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string){
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 先进后出 先Flush再Close

	pipeline.WriteSink(writer, p)

}

func createPipeline(fileName string,
	fileSize,
	chunkCount int) <-chan int {

	sortResults := []<-chan int{}

	chunkSize := fileSize / chunkCount // 必须保证整除，此处未做产品处理

	// 计时
	pipeline.Init()

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file),chunkSize)

		sortResults = append(sortResults,
			pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResults...)
}


func createNetworkPipeline(fileName string,
	fileSize,
	chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount // 必须保证整除，此处未做产品处理
	// 计时
	pipeline.Init()

	sortAddr := []string{}

	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i * chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file),chunkSize)

		//sortResults = append(sortResults,
		//	pipeline.InMemSort(source))
		addr := ":"+strconv.Itoa(7000 + i)

		pipeline.NetworkSink(addr,
			pipeline.InMemSort(source))

		sortAddr = append(sortAddr,addr)
	}

	sortResults := []<-chan int{}

	for _, addr := range sortAddr {
		sortResults = append(sortResults,
			pipeline.NetworkSource(addr))
	}

	return pipeline.MergeN(sortResults...)
}
