package main

import "fmt"

func main() {
	// 声明slice1并初始化
	var slice1 = []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)
	// 声明但是不分配空间
	var slice2 []int
	if slice2 == nil {
		fmt.Println("slice2 is empty")
		slice2 = make([]int, 3) //使用make分配空间
	}else{
		fmt.Println("slice2 is not empty")
	}
	slice2[0] = 1
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	// 声明并分配空间
	var slice3 []int = make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	slice4:= make([]int, 3)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

}