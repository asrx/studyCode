package main

import "fmt"

// 切片


func main() {
	
	nums := []int{2,3,4}

	sum := 0

	for _,num := range nums {
		sum += num
	}

	fmt.Println("sum:", sum)

	for i, num := range nums{
		if num == 3 {
			fmt.Println("index:",i)
		}
	}


	//range也可以用在map的键值对上。
	kvs := map[string]string{"a":"apple","b":"banana"}

	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	 //range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i,c)
	}

}


func printSlice(x []int) {

	fmt.Printf("len=%d cap=%d slice=%v\n", len(x),cap(x),x)
	
}
