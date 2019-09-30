package main

import (
	"fmt"
	b64 "encoding/base64"
)

// Base64编码

var p = fmt.Println

func main() {
	
	// 将要编解码的字符串
	data := "abc123!?$*&()'-=@~"

	// 编码需要 使用 []byte 类型的参数，所以要将字符串转成此类型
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	p(sEnc)

	// 解码可能会返回错误，如果不确定输入信息格式是否正确， 那么，你就需要进行错误检查了
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	p(string(sDec))
	p()

	// 使用 URL 兼容的 base64 格式进行编解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	p(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	p(string(uDec))

	// 标准 base64 编码和 URL 兼容 base64 编码的编码字符串存在 稍许不同（后缀为 + 和 -），但是两者都可以正确解码为 原始字符串
}

/*

*/
