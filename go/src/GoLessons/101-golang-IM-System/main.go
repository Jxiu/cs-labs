package main

//两个都是main包，所以不需要再引入了，可以一起编译
//go build -o server main.go server.go
// win:go build -o server.exe main.go server.go
// 简单测试：telnet ip port
func main(){
	server := NewServer("127.0.0.1", 9999)
	server.Start()
}