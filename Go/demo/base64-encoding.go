package main

import (
	"fmt"
	"encoding/base64"
)

func main() {

	data := "abc123!?$*&()'-=@~"

	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _:= base64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 使用 URL 兼容的base64 格式进行编码
	uEnc := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec,_:=base64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
