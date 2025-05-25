package main

import (
	"net"
	"fmt"
	"flag"
	"os"
	"io"
)

type Client struct{
	ServerIp string
	ServerPort int
	Name string
	conn net.Conn
	flag int //当前用户模式
}

func NewClient(serverIp string, serverPort int) *Client{
	// 创建客户端对象
	client := &Client{
		ServerIp: serverIp,
		ServerPort: serverPort,
		flag: 999,
	}
	// 连接server
	conn,err := net.Dial("tcp", fmt.Sprintf("%s:%d",serverIp, serverPort))
	if err != nil{
		fmt.Println("net.Dial error:",err)
		return nil
	}

	client.conn = conn
	return client
}

// 将服务器返回消息直接显示到标准输出
func (this *Client) DealResponse(){
	// 永久阻塞
	io.Copy(os.Stdout, this.conn)
/* 
	// 简写
	for {
		buf := make([]byte)
		this.conn.Read(buf)
		fmt.Println(buf)
	} */
}

// 菜单
func (this *Client) menu() bool{
	var flag int
	fmt.Println("1.共聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3{
		this.flag = flag
		return true
	}else{
		fmt.Println("功能不支持: " + string(flag))
		return false
	}
}

func (this *Client) SelectUser(){
	sendMsg := "who\n\n"
	_,err := this.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write error:",err)
		return
	}
}


func (this *Client) PrivateChat(){
	var remoteName string
	var chatMsg string
	this.SelectUser()
	fmt.Println(">>>>>>开始私聊，请输入聊天对象[用户名],exit退出：")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
			fmt.Println(">>>>>>请输入聊天内容，exit退出：")
			fmt.Scanln(&chatMsg)

			for chatMsg != "exit" {
				// 发送服务器
				if len(chatMsg) != 0 {
					sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
					_,err := this.conn.Write([]byte(sendMsg))
					if err != nil {
						fmt.Println("conn.Write error", err)
						break
					}
				}

				chatMsg = ""
				fmt.Println(">>>>>>请输入聊天内容，exit退出：")
				fmt.Scanln(&chatMsg)
			}

		remoteName = ""
		fmt.Println(">>>>>>请输入聊天内容,exit退出（重选聊天对象）：")
		fmt.Scanln(&remoteName)
	}
}

func (this *Client) PublicChat(){
	var chatMsg string
	fmt.Println(">>>>>>开始公聊，请输入聊天内容，exit退出：")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		// 发送服务器
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_,err := this.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn.Write error", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>>>请输入聊天内容，exit退出：")
		fmt.Scanln(&chatMsg)
	}

}

func (this *Client) UpdateName() bool {
	fmt.Println(">>>>>>请输入用户名：")
	fmt.Scanln(&this.Name)

	msg := "rename|" + this.Name +"\n"
	_,err := this.conn.Write([]byte(msg))
	if err != nil {
		fmt.Println("conn.Write error:", err)
		return false
	}
	return true
}
// 客户端主业务
func (this *Client) Run(){
	for this.flag != 0 {
		for this.menu() == false {

		}
		// 根据不同模式处理业务
		switch this.flag {
		case 1:
			// 公聊
			this.PublicChat()
			break
		case 2:
			// 私聊
			this.PrivateChat()
			break
		case 3:
			// 更新用户名
			this.UpdateName()
			break
		}

	}
}

// 解析命令行
var serverIp string
var serverPort int

// .\client -ip 127.0.0.1 -
// 解析命令行必须在main 之前
func init(){
	// 最后一个参数是输入 -h 时候的提示 如：.\client.exe -h
	flag.StringVar(&serverIp,"ip","127.0.0.1","设置服务器ip地址（默认是：127.0.0.1）")
	flag.IntVar(&serverPort,"port",9999 ,"设置服务器端口（默认是：9999）")
}

func main(){
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp , serverPort)
	if client == nil {
		fmt.Println(">>>>>>> connect to server error！！！")
		return
	}
	fmt.Println("connect to server success！！！")

	go client.DealResponse()

	// 启动客户端的业务
	client.Run()
}
