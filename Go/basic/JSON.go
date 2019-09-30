package main

import (
	"fmt"
	"encoding/json"
	"os"
)

// JSON
// Go 提供内置的 JSON 编解码支持，包括内置或者自定义类 型与 JSON 数据之间的转化

// http://blog.golang.org/2011/01/json-and-go.html
// http://golang.org/pkg/encoding/json/

type Response1 struct{
	Page int
	Fruits []string
}

type Response2 struct{
	Page int		`json:"pPage"`
	Fruits []string `json:"fFruits"`
}

var p = fmt.Println

func main() {

	// 基本数据类型到 JSON 字符串的编码 过程
	bolB, _ := json.Marshal(true)
	p(string(bolB))

	intB, _ := json.Marshal(1)
	p(string(intB))

	fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // JSON 包可以自动的编码你的自定义类型。编码仅输出可 导出的字段，并且默认使用他们的名字作为 JSON 数据的 键
    res1D := &Response1{
    	Page:	1,
    	Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    p(string(res1B))

    // Response2 的定义可以作为这个标签
    res2D := &Response2{
    	Page:	1,
    	Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    p(string(res2B))

    // 解码 JSON
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // 需要提供一个 JSON 包可以存放解码数据的变量。这里 的 map[string]interface{} 将保存一个 string 为键， 值为任意值的map
    var dat map[string]interface{}

    // 实际的解码和相关的错误检查
    if err := json.Unmarshal(byt, &dat); err != nil {
    	panic(err)
    }

    p(dat)

    // 将 num 的值转换成 float64 类型
    num := dat["num"].(float64)
    p(num)

    // 访问嵌套的值需要一系列的转化
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    p(str1)

    //
    str := `{"pPage": 1, "fFruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    p(res)
    p(res.Fruits[0])

    // 经常使用 byte 和 string 作为使用 
    // 标准输出时数据和 JSON 表示之间的中间值。
    // 也可以和 os.Stdout 一样，将 JSON 编码直接输出至 os.Writer 流中，
    // 或者作为 HTTP 响应体
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)

}


/*

*/
