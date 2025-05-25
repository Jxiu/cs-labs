package main

import "fmt"

// 类名首字母大写，表示该类是公共的，属性也是一样
type Hero struct {
	Name   string
	Level  int
	Attack int
}

/*
func (this Hero)Show(){
	fmt.Printf("Name: %s\n",this.Name)
	fmt.Printf("Level: %d\n",this.Level)
	fmt.Printf("Attack: %d\n",this.Attack)
}
func (this Hero) GetName() string{
	fmt.Printf("Name:%s\n",this.Name)
	return this.Name
}

func (this Hero) SetName(name string){
	// 值传递
	this.Name = name
}
*/

func (this *Hero) Show() {
	fmt.Printf("Name: %s\n", this.Name)
	fmt.Printf("Level: %d\n", this.Level)
	fmt.Printf("Attack: %d\n", this.Attack)
}
func (this *Hero) GetName() string {
	fmt.Printf("Name:%s\n", this.Name)
	return this.Name
}

func (this *Hero) SetName(name string) {
	this.Name = name
}
func main() {
	//结构体字面量,可以按照字段声明顺序，省略字段名称 如：hero := Hero{"张三",1,100}
	hero := Hero{Name: "张三", Level: 1, Attack: 100}
	hero.Show()
	hero.SetName("李4")
	hero.Show()
}
