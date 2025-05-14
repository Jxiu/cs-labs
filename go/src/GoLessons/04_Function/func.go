package main

import "fmt"

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