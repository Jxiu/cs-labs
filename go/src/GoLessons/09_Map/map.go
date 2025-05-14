package main

import "fmt"
func main(){

	// =====> 第一种声明方式
	// 声明一个map，key为string，value为string
	var myMap1 map [string]string
	if(myMap1 == nil){
		fmt.Println("myMap1 is nil")
	}
	// 使用make分配空间
	myMap1 = make(map[string]string,10)
	myMap1["one"] = "java"
	myMap1["two"] = "javascript"
	myMap1["three"] = "python"

	fmt.Println(myMap1)

	// =====> 第二种声明方式，声明默认容量的map
	myMap2 := make(map[int]string)
	myMap2[1] = "java"
	myMap2[2] = "javascript"
	myMap2[3] = "python"
	fmt.Println(myMap2)

	// =====> 第三种声明方式，声明默认容量的map,注意多个键值对使用逗号分隔,且最后一个后面也必须有逗号
	myMap3 := map[string]float32{
		"java":1.2,
		"javascript":2.3,
		"python":3.4,
	}
	fmt.Println(myMap3)
}