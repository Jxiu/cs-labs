package main

import (
	"fmt"
)

func printArray(arr []int){
	// 切片是引用传递
	// _ 表示匿名变量，可以只声明不使用
	for _,v :=range arr{
		fmt.Println("value = ", v)
	}
	arr[0] = -1
}

func main(){
	myarr := []int{1,2,3,4} // 动态数组，切片 slice
	fmt.Printf("myarr tyeps is %T\n", myarr)
	printArray(myarr)
	printArray(myarr)
}