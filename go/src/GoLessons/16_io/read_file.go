package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readTxt() {
	// 相对于当前可执行文件的路径
	var realitivePath = "a.txt"
	byteDta, err := os.ReadFile(realitivePath)
	if err != nil {
		fmt.Println("os.ReadFile error:", err)
		return
	}
	fmt.Println(string(byteDta))
}

func readBigFile() {
	//
	file, err := os.Open("big.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	for {
		var bytes = make([]byte, 1024)
		n, err := file.Read(bytes)
		if err == io.EOF {
			break
		}
		fmt.Println(string(bytes), "\n byte length:", n)
	}
}

func readBuffer() {
	file, err := os.Open("big.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	for {
		line, _, e := buf.ReadLine()
		if e == io.EOF {
			break
		}
		fmt.Printf(string(line))
	}
	fmt.Println()
}

func readByWorld() {
	file, err := os.Open("big.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	// lamda
	buf.Split(bufio.ScanWords)
	var index = 0
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}
}
func main() {
	// readTxt()
	// readBigFile()
	// readBuffer()
	readByWorld()
}
