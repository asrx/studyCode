package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
	"time"
	"fmt"
)

var startTime time.Time

func Init(){
	startTime = time.Now()
}

func ArraySource(a ...int) <-chan int {

	out := make(chan int)
	go func() {
		for _, v:= range a {
			out <- v
		}
		close(out)
	}()
	return out

}

func InMemSort(in <-chan int) <-chan int {

	out := make(chan int, 1024)

	go func() {

		// Read into memory
		a := []int{}
		for v := range in {
			a = append(a,v)
		}

		fmt.Println("Read done",
			time.Now().Sub(startTime))

		// Sort
		sort.Ints(a)

		fmt.Println("InMemSort done",
			time.Now().Sub(startTime))

		// Output
		for _, v := range a {
			out <- v
		}
		close(out)

		fmt.Println("Merge done",
			time.Now().Sub(startTime))
	}()

	return out
}

// 归并节点
func Merge(in1, in2 <-chan int) <-chan int  {

	out := make(chan int, 1024)

	go func() {
		v1, ok1 := <- in1
		v2, ok2 := <- in2
		
		for ok1 || ok2  {// 当ok1，ok2 其中一个有数据时

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

// 读 - 支持分块
func ReaderSource(reader io.Reader, chunkSize int) <-chan int{
	out := make(chan int, 1024)

	go func() {
		buffer := make([]byte, 8)
		bytesRead := 0
		for {
			// n 表示读了多少个字节
			n, err := reader.Read(buffer)
			bytesRead += n
			if n > 0 {
				v := int(
					binary.BigEndian.Uint64(buffer))
				out <- v
			}

			// -1 为全部读
			if err != nil ||
				(chunkSize != -1 &&
					bytesRead >= chunkSize){
				break
			}
		}
		close(out)
	}()

	return out
}

// 写
func WriteSink(writer io.Writer, in <-chan int) {
	for v := range in {
		buffer := make([]byte, 8)
		binary.BigEndian.PutUint64(
			buffer,
			uint64(v))

		writer.Write(buffer)
	}
}
// 随机数据源
func RandomSource(count int) <-chan int{
	out := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()

	return out
}

// 两两归并 - 搭建归并结点组
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