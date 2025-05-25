package main

import "fmt"

// 声明一种数据类型 myint 是int一个别名
type myint int

type Book struct {
	title  string
	author string
}

func cahegBook(book Book) {
	// 传递的是book的副本
	book.title = "python"

}

// 使用指针修改属性，可以直接点，但是获取变量需要使用 *book
func changeBook(book *Book) {
	book.author = "lisi666"
	fmt.Printf("---> %v\n", book)
	fmt.Printf("---> %v\n", *book)
}
func main() {

	var book1 Book
	book1.title = "golang"
	book1.author = "zhangsan"

	fmt.Printf("%v, %T\n", book1, book1)

	cahegBook(book1)
	fmt.Printf("%v\n", book1)

	changeBook(&book1)
	fmt.Printf("%v\n", &book1)

}
