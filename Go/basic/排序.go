package main

import (
	"fmt"
	"sort"
)

// 排序

func main() {

	strs := []string{"c","a","b"}

	sort.Strings(strs)

	fmt.Println("String:", strs)

	ints := []int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints:	", ints)

	// 检查一个序列是不是已经 是排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted:", s)
	

}

/*

*/
