package main

/*
	声明变量的四种方式
*/
import (
	"fmt"
)

// 声明全局变量 方法1、2、3是可以使用的
var NONE int
var A float32 = 3.14
var B int64 = 123456789
var C string = "hello world"
var CC string = `aa
 aa\n 多行
 字符串`
var D bool = true
var E byte = 'a'         // byte类型 int8
var F rune = '中'         //unicode编码 int32
var G complex64 = 1 + 2i //复数
var H uint8 = 255        //无符号数

var G_A = 1

// 第四种方法,不能声明全局变量
// G_4 := 1
func main() {
	// 1.声明一个变量默认值为0
	var a int
	fmt.Println("a =", a)
	fmt.Printf("Type of a = %T\n", a)
	// 2.声明变量并给一个初始化值
	var b int = 1
	fmt.Println("b =", b)
	fmt.Printf("Type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Println("bb = %s, type of bb = %T\n", bb, bb)
	// 3.声明变量，并使用类型推断 不是很推荐
	var c = 2
	fmt.Println("c =", c)
	fmt.Printf("Type of c = %T\n", c)

	var cc = "abcd"
	fmt.Println("cc = %s, type of cc = %T\n", cc, cc)
	// 4.省去var关键字，初始化赋值并直接推断类型，开发中最常使用
	d := 3
	fmt.Println("d =", d)
	fmt.Printf("Type of d = %T\n", d)

	dd := "mm"
	fmt.Printf("dd = %s, type of dd = %T\n", dd, dd)

	f := 3.14
	fmt.Printf("f = %f, type of f = %T\n", f, f)

	// 全局变量的使用
	fmt.Println("A =", A)
	fmt.Println("B =", B)
	fmt.Println("C =", C)
	fmt.Println("CC =", CC)
	fmt.Println("D =", D)
	fmt.Println("E =", E)
	fmt.Printf("F = %s\n", F)
	fmt.Println("G =", G)
	fmt.Println("H =", H)
	fmt.Println("G_A =", G_A)
	fmt.Println("NONE =", NONE)

	// fmt.Println("G_4 =", G_4)

	//--------------------
	// 声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx =", xx, ", yy =", yy)
	var kk, ll = 11, "goslin"
	fmt.Println("kk =", kk, ", ll =", ll)

	var (
		vv        int  = 100
		isSuccess bool = true
	)
	fmt.Println("vv =", vv, ", isSuccess =", isSuccess)
}
