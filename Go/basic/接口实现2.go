package main

import (
	"fmt"
)

type Man interface{
	name() string
	age() int
}

type Woman struct{}

func (woman Woman) name() string {
	return "Jin Yawei"
}

func (woman Woman) age() int {
	return 23
}

type Men struct{}

func (men Men) name() string {
	return "Donovan.Xu"
}

func (men Men) age() int {
	return 32
}

func main() {
	
	var man Man

	man = new(Woman)
	fmt.Println(man.name())
	fmt.Println(man.age())

	man = new(Men)
	fmt.Println(man.name())
	fmt.Println(man.age())	

}


func (name string) imp() string{
    print("这是实现方法的写法")
}

func sum(x int,y int) int{
    print("这是正常写法")
}


