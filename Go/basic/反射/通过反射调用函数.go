package main

import (
	"reflect"
	"fmt"
)

// 普通函数
func add(a, b int) int{
	return a + b
}

func main() {

	// 将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)

	// 构造函数参数，传入两个整型值
	// 将 10 和 20 两个整型值使用 reflect.ValueOf 包装为 reflect.Value，再将反射值对象的切片 []reflect.Value 作为函数的参数
	paramList := []reflect.Value{
		reflect.ValueOf(10),
		reflect.ValueOf(20),
	}

	// 反射调用函数
	// 使用 funcValue 函数值对象的 Call() 方法，传入参数列表 paramList 调用 add() 函数
	retList := funcValue.Call(paramList)

	// 获取第一个返回值, 取整数值
	// 调用成功后，通过 retList[0] 取返回值的第一个参数，使用 Int 取返回值的整数值
	fmt.Println(retList[0].Int())

}



