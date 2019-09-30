package main

import (
	"reflect"
	"fmt"
)

func main() {

	var a int = 1024

	// 获取变量a的反射值对象
	//valueOfA := reflect.ValueOf(a)
	// 尝试将a修改为1(此处会发生崩溃)
	// panic: reflect: reflect.Value.SetInt using unaddressable value


	/*
	使用 reflect.Value 类型的 Elem() 方法获取 a 地址的元素，
	也就是 a 的值。reflect.Value 的 Elem() 方法返回的值类型也是 reflect.Value
	*/
	// 获取变量a的反射值对象(a的地址)
	valueOfA := reflect.ValueOf(&a)

	// 取出a地址的元素(a的值)
	valueOfA = valueOfA.Elem()

	valueOfA.SetInt(1)

	fmt.Println(valueOfA.Int())
	fmt.Println("*****************")

	type dog struct {
		LegCount int
	}
	// 获取dog实例地址的反射值对象
	valueOfDog := reflect.ValueOf(&dog{})

	// 取出dog实例地址的元素
	valueOfDog = valueOfDog.Elem()

	// 获取legCount字段的值
	vLegCount := valueOfDog.FieldByName("LegCount")

	// 尝试设置legCount的值
	vLegCount.SetInt(4)

	fmt.Println(vLegCount.Int())

}



