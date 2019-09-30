package main

import (
	"fmt"
	"reflect"
)

// IsNil() 常被用于判断指针是否为空；IsValid() 常被用于判定返回值是否有效
func main() {

	// *int的空指针
	var a * int

	// 将变量 a 包装为 reflect.Value 并且判断是否为空，此时变量 a 为空指针，因此返回 true。
	fmt.Println("var a *int:", reflect.ValueOf(a).IsNil())

	// nil值
	// 对 nil 进行 IsValid() 判定（有效性判定），返回 false
	fmt.Println("nil:", reflect.ValueOf(nil).IsValid())

	// *int类型的空指针
	fmt.Print("(*int)(nil):",reflect.ValueOf((*int)(nil)).Elem().IsValid())

	s := struct {
	}{}

	// 尝试从结构体中查找一个不存在的字段
	// 通过 FieldByName 查找 s 结构体中一个空字符串的成员，如成员不存在，IsValid() 返回 false
	fmt.Println("不存在的结构体成员:",
		reflect.ValueOf(s).FieldByName("").IsValid())

	// 尝试从结构体中查找一个不存在的方法
	// 通过 MethodByName 查找 s 结构体中一个空字符串的方法，如方法不存在，IsValid() 返回 false
	fmt.Println("不存在的结构体方法:",
		reflect.ValueOf(s).MethodByName("").IsValid())

	m := map[int]int{}

	//MapIndex() 方法能根据给定的 reflect.Value 类型的值查找 map，并且返回查找到的结果
	fmt.Println("不存在的键:",
		reflect.ValueOf(m).MapIndex(reflect.ValueOf(3)).IsValid())
}



