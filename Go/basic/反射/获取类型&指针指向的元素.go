package main

import (
	"reflect"
	"fmt"
)

type Enum int

const (
	Zero Enum = 0
)

func main() {

	var a int

	 typeOfA := reflect.TypeOf(a)

	 fmt.Println(typeOfA.Name(), typeOfA.Kind())

	 fmt.Println("******************")
	// 声明一个空结构体
	type cat struct {
	}

	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(cat{})

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())

	// 获取Zero常量的反射类型对象
	typeOfA = reflect.TypeOf(Zero)

	// 显示反射类型对象的名称和种类
	fmt.Println(typeOfA.Name(), typeOfA.Kind())

	fmt.Println("******************")
	// 创建cat的实例
	ins:= &cat{}

	// 获取结构体实例的反射类型对象
	typeOfCat = reflect.TypeOf(ins)

	// 显示反射类型对象的名称和种类
	fmt.Printf("name: '%v' kind:'%v'\n",
		typeOfCat.Name(),typeOfCat.Kind())

	fmt.Println("******************")
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()

	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n",
		typeOfCat.Name(), typeOfCat.Kind())


}



