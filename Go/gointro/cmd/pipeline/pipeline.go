package main

import (
	"sort"
	"fmt"
	"math/rand"
	"io"
	"encoding/binary"
	"os"
	"bufio"
	"time"
)

var startTime time.Time

func Init(){
	startTime = time.Now()
}

// 加载原始数据（数组）到管道
func ArraySource(a ...int) <-chan int{
	out := make(chan int)

	// channgel 是 goroutine 与 goroutine 之间的通信，不能直接导，必须起一个goroutine
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

// 管道中的数据加载到内存并排序
func InMemSort(in <-chan int) <-chan int{
	out := make(chan int)

	go func() {
		arr := []int{}
		// channgel 中的数据导入到内存 arr 中
		for v := range in {
			arr = append(arr,v)
		}

		fmt.Println("Read done ",
			time.Now().Sub(startTime))

		// 内置排序函数
		sort.Ints(arr)

		fmt.Println("InMem done ",
			time.Now().Sub(startTime))

		// 重新载入channel
		for _, v := range arr {
			out <- v
		}
		close(out)

		fmt.Println("Merge done ",
			time.Now().Sub(startTime))
	}()

	return out
}

// 归并节点
func Merge(in1 , in2 <-chan int,) <-chan int{
	out := make(chan int)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {	// 当ok1，ok2 其中一个有数据时

			// 每次送出一条数据【v1 or v2】

			/*
				1. 当ok2没数据，那么输出v1
				2. 或当ok2有数据，需要保证 in1 有数据，并且v1<=v2
			*/
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				v1, ok1 = <-in1
			}else{
				out <- v2
				v2, ok2 = <-in2
			}
		}

		close(out)
	}()

	return out
}

/*************************************************/
// 生成 count 个随机 int
func RandomSource(count int) <-chan int{
	out := make(chan int,1024)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer, in <-chan int){
	for v:=range in{

		buffer := make([]byte,8)

		binary.BigEndian.PutUint64(buffer,
			uint64(v))

		writer.Write(buffer)
	}
}

func ReadSource(reader io.Reader,
	chunkSize int) <-chan int{

	out := make(chan int, 1024)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			// n 读了多少个字节
			n, err := reader.Read(buffer)
			bytesRead += n

			if n > 0 {
				// 当读到数据时，把buffer
				v := int(binary.BigEndian.Uint64(buffer))

				out <- v
			}

			// -1 全部读取，限制读取字节数 chunkSize
			if err != nil ||
				(chunkSize != -1 &&
					bytesRead >=  chunkSize){
				break
			}
		}
		close(out)
	}()

	return out
}

// 多路，两两归并 - 搭建归并结点组
func MergeN(inputs ...<-chan int) <-chan int{

	if len(inputs) == 1 {
		return inputs[0]
	}

	m := len(inputs) / 2

	// merge inputs[0..m] and inputs [m..end]
	return Merge(
		MergeN(inputs[:m]...),
		MergeN(inputs[m:]...))

}


func createPipeline(fileName string,
	fileSize,chunkCount int) <-chan int{

	// 搜集到的结果
	sortResults := []<-chan int{}

	// 每块的大小（必须保证整除，此处未做产品处理）
	chunkSize := fileSize / chunkCount

	// 计时
	Init()

	// 原始数据分块读取
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		// 移动到开始读的位置（从第几块开始，0 为从头开始）
		file.Seek(int64(i*chunkSize),0)

		source := ReadSource(
			bufio.NewReader(file),chunkSize)

		// 搜集到内存中
		sortResults = append(sortResults,
			InMemSort(source))
	}

	return MergeN(sortResults...)
}

// 排序后写入文件
func writeToFile(p <-chan int,fileName string){
	// 创建文件
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 写入
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	WriteSink(writer,p)
}

// 打印排序后的文件
func printFile(fileName string){
	file, err := os.Open(fileName)
	if err !=nil {
		panic(err)
	}
	defer file.Close()

	p := ReadSource(bufio.NewReader(file),-1)
	count := 0

	for v := range p{
		fmt.Println(v)
		if count++; count >= 100 {
			break
		}
	}
}

func mergePipeline(){
	const fileName = "small.in"
	const outFileName = "small.out"
	const fileSize = 512

	//const fileName = "large.in"
	//const outFileName = "large.out"
	//const fileSize = 800000000

	const chunkCount = 4

	// 创建
	p := createPipeline(fileName,fileSize,chunkCount)

	// 排序后写入文件
	writeToFile(p,outFileName)

	// 打印排序后的文件
	printFile(outFileName)
}

// 生成测试数据
func mergeDemoSource(){

	const fileName = "small.in"
	const n = 64

	//const fileName = "large.in"
	//const n = 100000000

	// 写入数据
	file, err := os.Create(fileName)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	// 生成随机数据源
	p := RandomSource(n)
	writer := bufio.NewWriter(file)
	WriteSink(writer, p)
	writer.Flush()

	// 读取数据 - 打印前100 （测试）
	file, err = os.Open(fileName)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	p = ReadSource(bufio.NewReader(file), -1)
	count := 0
	for v:=range p{
		fmt.Println(v)
		count++
		if count > 100{
			break
		}
	}
}

//
func mergeDemo() {
	p := Merge(InMemSort(
		ArraySource(3,2,6,9,4)),
		InMemSort(
			ArraySource(7,0,5,1,13,8)))

	for v := range p {
		fmt.Println(v)
	}
}


func main() {

	// mergeDemo()

	// 生成测试数据
	// mergeDemoSource()

	//
	mergePipeline()

}
