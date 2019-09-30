package main

import (
	"reflect"
	"fmt"
)

type dummy struct {
	a int
	b string

	float32
	bool

	next *dummy
}

func main() {

	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	// 获取字段数量
	fmt.Println("NumField:", d.NumField())

	// 获取索引为2的字段(float32字段)
	floatField := d.Field(2)

	fmt.Println("Field:",floatField.Type())

	// 根据名字查找字段
	// []int{4,0} 中的 4 表示，在 dummy 结构中索引值为 4 的成员，
	// 也就是 next。next 的类型为 dummy，也是一个结构体，
	// 因此使用 []int{4,0} 中的 0 继续在 next 值的基础上索引，结构为 dummy 中索引值为 0 的 a 字段，类型为 int
	fmt.Println("FieldByIndex([]int{4,0}).Type():",
		d.FieldByIndex([]int{4,0}).Type())

	fmt.Println("FieldByIndex([]int{4,0}).Type():",
		d.FieldByIndex([]int{1}).Type())
}



