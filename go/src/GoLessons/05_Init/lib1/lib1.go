package lib1 //通常包名和文件名一致

import "fmt"
// 公开的函数需要首字母大写
func Lib1Test() {
	fmt.Println("lib1 test()...")
}
func init() {
	fmt.Println("lib1 init()...")
}