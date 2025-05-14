package main

import "fmt"
/* 
//值传递，不能正确交换主函数中的变量
func swap(a,b int){
	var temp int
	temp = a
	a = b 
	b = temp
}
 */

func swap(pa *int, pb *int){
	var temp int;
	temp = *pa;
	*pa = *pb;
	*pb = temp;

}
func main(){
	var a int = 10 // 必须明确定义类型，否则  cannot use b (type int) as type *int in argument to swap
	var b int = 20

	swap(&a, &b)
	fmt.Println("a = ", a, ", b = ", b)

	var p *int;
	p = &a;
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	// 二级指针
	var pp **int
	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(**pp)

}