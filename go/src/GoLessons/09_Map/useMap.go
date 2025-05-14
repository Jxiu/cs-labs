package main

import "fmt"

func printMap(cityMap map[string]string) {
	// 引用传递
	for k,v := range cityMap{
		fmt.Println(k,v)
	}
}

func addCity(cityMap map[string]string){
	cityMap["UK"] = "伦敦"
}
func main() {
	cityMap := make(map[string]string)

	cityMap["China"] = "北京"
	cityMap["Japan"] = "东京"
	cityMap["USA"] = "纽约"

	for k,v := range cityMap {
		fmt.Println(k,v)
	}
	// 删除 delete、copy、make 都是关键字
	delete(cityMap, "Japan")

	cityMap["USA"] = "las vegas"

	fmt.Println("-----------")
	addCity(cityMap)
	printMap(cityMap)
}