package main

import "fmt"
/**
 * 方法的定义规则
接收者类型：方法的接收者类型可以是任何自定义类型（结构体、整数、字符串等）。
方法名：方法名必须是唯一的，但在不同的类型上可以有相同的方法名（这称为方法重载，但 Go 中的方法重载与传统意义上的重载有所不同）。
参数和返回值：方法可以有多个参数和返回值，与普通函数类似。

	类型：
	func [(receiverType)] [funcName]([paramType paramName, paramType paramName]) [(returnType,returnType)] {}

	func(){} 匿名函数
 */

func foo1(a string, b int) int {
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	c := 333
	return c
}

// 多个返回值,匿名
func foo2(a string, b int) (int, int){
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	c := 100
	d := 200
	return c, d
}
// 多个返回值，命名
func foo3(a string, b int) (r1 int, r2 int){
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	fmt.Println("r1= ", r1, "r2= ", r2)
	// 给命名返回值变量赋值
	r1 = 1000
	r2 = 3000
	return
}
// 命名返回值，同类型
func foo4(a string, b int) (r1,r2 float32){
	fmt.Println("a= ", a)
	fmt.Println("b= ", b)

	// 给命名返回值变量赋值
	r1 = 1000.0
	r2 = 2000.0
	return
}
func main() {

	c := foo1("abc", 555)
	fmt.Println("c= ", c)

	cc, dd := foo2("foo2", 999)
	fmt.Println("cc= ", cc, "dd= ", dd)

	m, n := foo3("foo3", 666)
	fmt.Println("m= ", m, "n= ", n)

	x, y := foo4("foo4", 777)
	fmt.Println("x= ", x, "y= ", y)
}