package main

import (
	"fmt"
	"os"
	"../../pipeline"
	"bufio"
)

func main() {
	//const filename = "small.in"
	//const n = 64

	const filename = "large.in"
	const n = 100000000

	file, err := os.Create(filename)
	defer file.Close()

	if err != nil {
		panic(err)
	}

	p := pipeline.RandomSource(n)
	//pipeline.WriteSink(file,p)
	writer := bufio.NewWriter(file)
	pipeline.WriteSink(
		writer,p)
	writer.Flush()


	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(
		bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count ++
		if count >= 100{
			break
		}
	}

	//mergeDemo()

}

func mergeDemo() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3,2,6,7,4)),
		pipeline.InMemSort(
			pipeline.ArraySource(7,4,0,3,2,13,8)))

	for v := range p{
		fmt.Println(v)
	}
}
