package main

import (
	"fmt"
	"io"
	"os"
)

func write() {
	// 创建文件，写入，第二个参数是文件打开模式类似java
	file, e := os.OpenFile("w.txt", os.O_CREATE|os.O_WRONLY, 0777)
	if e != nil {
		panic(e)
	}
	defer file.Close()
	// file.WriteString("aaaaaaaaa")
	/*
		// 不能读 io.ReadAll err:  read w.txt: Access is denied
		byteData, e := io.ReadAll(file)
		if e != nil {
			fmt.Println("io.ReadAll err: ", e)
			return
		}
		fmt.Println(string(byteData))
	*/
}

func copy() {
	source, err := os.Open("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer source.Close()
	target, err := os.OpenFile("copy.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer target.Close()
	io.Copy(target, source)
}

func main() {
	// write()
	copy()
}
