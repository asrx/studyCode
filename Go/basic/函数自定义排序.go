package main

import (
	"fmt"
	"sort"
)

// 函数自定义排序

// 创建一个为内置 []string 类型
type ByLength []string

// 在类型中实现了 sort.Interface 的 Len，Less 和 Swap 方法
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
	

}

/*

*/
