package main // 包名,如果不是main则不能使用go run 运行否则报错：go run: cannot run non-main package
// 导入包
/* 
import "fmt"
import "time" 
*/
import (
	"fmt"
	"time"
)

// 运行：go run hello.go
// 编译成二进制执行文件：go build hello.go
// main 函数
func main()  { //左大括号 必须与方法名在同一行
	// golang 中表达式结尾，加';' 与不加相同
	time.Sleep(1 * time.Second);
	fmt.Print("hello world!")
}
