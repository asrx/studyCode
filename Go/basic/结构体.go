package main

import "fmt"

// 语言结构体

type Books struct {
	title string
	author string
	subject string
	book_id int
}

func main() {
	
	// 创建一个新的结构体
	// fmt.Println(Books{"Go 语言","www.runoob.com","Go 语言教程",649507})

	// 也可以使用 Key => value 格式
	// fmt.Println(Books{title:"Go 语言",author:"runoob.com",subject:"Go 语言教程",book_id:649507})	

	// 忽略的字段为 0 或 空
	// fmt.Println(Books{title:"Go 语言", author:"runoob.com"})

	var Book1 Books
	var Book2 Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(Book1)

	printBook(Book2)

}

func printBook( book Books){
	fmt.Printf("Book title : %s\n",book.title)
	fmt.Printf("Book author : %s\n",book.author)
	fmt.Printf("Book subject : %s\n",book.subject)
	fmt.Printf("Book book_id : %d\n",book.book_id)

	fmt.Printf("\n")
}

