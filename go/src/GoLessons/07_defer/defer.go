package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer func call...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func call..")
	return 0
}

// return 函数先执行
func deferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

func main(){
	// 写入defer关键 跟一个表达式或函数调用,类似java中的finally
	defer fmt.Println("main end 1")
	// 按顺序压栈，逆序执行
	defer fmt.Println("main end 2")
	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
	deferAndReturn()
}