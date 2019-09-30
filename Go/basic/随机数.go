package main

import (
	"fmt"
	"time"
	"math/rand"
)

// 随机数
// http://golang.org/pkg/math/rand/

var p = fmt.Println

func main() {

	// rand.Intn 返回一个随机的整数 n，0 <= n <= 100
	p(rand.Intn(100), ",")
	p(rand.Intn(100))
	p()

	// rand.Float64 返回一个64位浮点数 f， 0.0 <= f <= 1.0
	p(rand.Float64())

	// 生成其他范围的随机浮点数，例如 5.0 <= f <= 10.0
	p((rand.Float64() * 5 ) + 5, ",")
	p((rand.Float64() * 5 ) + 5)
	p()

	// 
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	p(r1.Intn(100), ",")
	p(r1.Intn(100))
	p()

	// 如果使用相同的种子生成的随机数生成器，将会产生相同的随机 数序列
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	p(r2.Intn(100), ",")
	p(r2.Intn(100))
	p()

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	p(r3.Intn(100), ",")
	p(r3.Intn(100))





}


/*

*/
