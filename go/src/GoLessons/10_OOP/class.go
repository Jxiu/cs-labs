package main

import "fmt"

type Hero struct {
	Name string
	Level int
	Attack int
}

func (this Hero) GetName(){
	fmt.Printf("name:%s\n",this.Name)
}

func (this Hero) SetName(name string){
	this.Name = name
}

func main(){
	hero := Hero{Name:"张三",Level:1,Attack:100}
}