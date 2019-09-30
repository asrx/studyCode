package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()

	l.PushBack("first")
	l.PushFront(67)
	//l2 := list.List{}

	for i := l.Front(); i!=nil;i=i.Next()  {
		fmt.Println(i.Value)
	}
	//fmt.Println(reflect.TypeOf(l))
}
