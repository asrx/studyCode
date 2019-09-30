package main

import (
	"fmt"
	"crypto/sha1"
)

// SHA1 散列经常用 生成二进制文件或者文本块的短标识

var p = fmt.Println

func main() {
	
	s := "sha1 this string"

	// 产生一个散列值的方式是 sha1.New()，sha1.Write(bytes)， 然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New()

	// 写入要处理的字节。如果是一个字符串，需要使用 []byte(s) 来强制转换成字节数组
	h.Write([]byte(s))

	// 用来得到最终的散列值的字符切片。Sum 的参数可以 用来给现有的字符切片追加额外的字节切片：一般不需要
	bs := h.Sum(nil)

	p(s)
	// SHA1 值经常以 16 进制输出，例如在 git commit 中。使用 %x 来将散列结果格式化为 16 进制字符串
	fmt.Printf("%x\n",bs)

	/*
	你可以使用和上面相似的方式来计算其他形式的散列值。例 如，计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法
	*/

}

/*
运行程序计算散列值并以可读 16 进制格式输出
*/
