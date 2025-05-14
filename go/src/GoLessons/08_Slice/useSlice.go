package main

import "fmt"

func main(){
	// 声明一个slice 长度3 容量5（分配空间，但是不可访问）
	var numbers = make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 向numbers 追加两个元素1,2
	numbers = append(numbers, 1, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	// 容量满了，在添加元素会自动扩容，为原来的2倍
	numbers = append(numbers,3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []int{1,2,3}
	s1 := s[0:2] // 类似java的subList
	fmt.Println(s1)
	s1[0] = 100
	fmt.Println(s1, s)

	// copy 可以将slice深拷贝
	s2 := make([]int, 3)
	copy(s2, s)
	fmt.Println(s2)

}