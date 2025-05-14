package main 

import "fmt"
const (
	// 声明枚举
	// 可以在coonst() 中添加一个关键字 iota,每行的值都会累加1,第一行的iota = 0
	MONDAY = iota //iota = 0
	TUESDAY       //iota = 1
	WEDNESDAY	  //iota = 2
	THURSDAY	  //iota = 3
	FRIDAY		  //iota = 4
	SATURDAY      //iota = 5
	SUNDAY	      //iota = 6
)

const(
	A = 10 * iota
	B
	C
)

const (
	a,b = iota+1, iota+2 // 声明公式复制
	c,d
	e,f

	g,h = iota*2, iota*3 // 公式改变，不影响iota逐行递增
	i,k
)
func main() {
	const length int = 10
	// length = 20 变量只读，不能修改
	fmt.Println("length = ", length)

	fmt.Println("MONDAY = ", MONDAY)
	fmt.Println("TUESDAY = ", TUESDAY)
	fmt.Println("WEDNESDAY = ", WEDNESDAY)
	fmt.Println("THURSDAY = ", THURSDAY)
	fmt.Println("FRIDAY = ", FRIDAY)
	fmt.Println("SATURDAY = ", SATURDAY)
	fmt.Println("SUNDAY = ", SUNDAY)

	fmt.Println("-------------------- ")
	fmt.Println("A = ", A)
	fmt.Println("B = ", B)
	fmt.Println("C = ", C)
	fmt.Println("-------------------- ")
	fmt.Println("a = ", a, ", b = ", b)
	fmt.Println("c = ", c, ", d = ", d)
	fmt.Println("e = ", e, ", f = ", f)
	fmt.Println("g = ", g, ", h = ", h)
	fmt.Println("i = ", i, ", k = ", k)

	// var z int = iota //undefined: iota 只能在常量块中定义
}