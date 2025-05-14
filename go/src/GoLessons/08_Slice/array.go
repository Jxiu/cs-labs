package main

import "fmt"

// 静态数据组参数，只能接收固定长度的数组
// cannot use myArr (type [10]int) as type [4]int in argument to printArray
func printArrayLen4(arr [4]int){
	// 传值方式是值拷贝
	for i,v := range arr{
		fmt.Println("index = ", i, ",value = ", v)
	}
	arr[0] = -11
}
func main() {
	var myArr [10]int //固定长度数组
	myArr2 := [10]int{1,2,3,4}
	myArr3 := [4]int{11,22,33,44}
	// for i:=0; i<10; i++{
	for i:=0; i<len(myArr); i++{
		fmt.Println(myArr[i])
	} 

	for index,value := range myArr2 {
		fmt.Println("index = ", index, ",value = ", value)
	}

	fmt.Printf("myArr  types= %T\n",myArr)
	fmt.Printf("myArr2 types= %T\n",myArr2)
	fmt.Printf("myArr3 types= %T\n",myArr3)

	printArrayLen4(myArr3)

	for i,v := range myArr3{
		fmt.Println("index = ", i, ",value = ", v)
	}

}