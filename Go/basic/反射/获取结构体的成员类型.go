package main

import (
	"reflect"
	"fmt"
)

func main() {

	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	ins := cat{Name: "mimi", Type: 1}

	typeOfCat := reflect.TypeOf(ins)

	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)

		//
		fmt.Printf("name: %v tag: '%v'\n",
			fieldType.Name,
			fieldType.Tag)

	}

	// 通过字段名，找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type");ok{

		fmt.Println(catType.Tag.Get("json"),
			catType.Tag.Get("id"))
	}
}



