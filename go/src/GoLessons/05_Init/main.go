package main
// 执行到import的包时，会先执行init()函数，递归执行import的包，直到执行到main.go
import (
	// 匿名导入，可以只调用init方法，不使用其他函数（无法使用），不会编译错误
	_ "GoLessons/05_Init/lib1" // 相对于$GOROOT或$GOPATH的路径
	// localLib2 "GoLessons/05_Init/lib2" // 别名
	. "GoLessons/05_Init/lib2" // 类似java import static
)
func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()
	// 使用包别名调用
	// localLib2.Lib2Test()
	Lib2Test()
}