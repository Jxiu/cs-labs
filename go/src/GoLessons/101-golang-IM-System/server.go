package main

import (
	"net"
	"fmt"
	"sync"
	"io"
	"time"
)

type Server struct{
	Ip string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	// 锁
	mapLock sync.RWMutex
	// 广播channel
	MsgChan chan string
}
// server构造函数
func NewServer(ip string, port int) *Server{
	server := &Server{
		Ip: ip,
		Port: port, // 注意最后的，不能省略
		OnlineMap: make(map[string]*User),
		MsgChan: make(chan string),
	}
	return server
}
// 监听广播channel的goroutine，一旦有消息就发送给所有用户
func (this *Server) ListenMsg(){
	for{
		msg := <-this.MsgChan

		this.mapLock.Lock()
		for _,user := range this.OnlineMap {
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}
func (this *Server) BroadCast(user *User, msg string){
	sendMsg := "["+user.Addr+"]" + user.Name + ": " + msg
	// 发送到广播队列
	this.MsgChan <- sendMsg
}
func (this *Server)Handler(conn net.Conn){
	//处理已经连接的业务
	fmt.Println("连接建立成功...")

	// 创建用户
	user := NewUser(conn, this)
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	// 接收用户发送的消息
	go func(){
		buf := make([]byte, 4096)
		for{
			n,err := conn.Read(buf)
			if n == 0{
				user.Offline()
				return
			}
			if err != nil && err != io.EOF{
				// 读取异常
				fmt.Println("conn Read error:", err)
			}
			// 提取用户消息，去除\n
			msg := string(buf[:n-1])
			// 将得到的消息广播
			user.DoMsg(msg)
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for{
		select {
		case <- isLive:
			//当前用户活跃重置定时器
			// 不做任何处理，为了激活select，更新下面定时器
			
			//定时器，结果是channel,尝试执行条件，不满足条件会重置定时器
		case <- time.After(time.Second * 180):
			//已经超时
			// 将当前用户强制关闭
			user.SendMsg("you are OFFLINE!!")
			// 销毁用户资源
			close(user.C)
			conn.Close()

			return //runtime.Goexit()
		}
	}
	
}
// 启动,成员方法，方法中有指针接收者
func (this *Server)Start(){
	// 监听端口
	listener,err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen error:",err)
		return
	}
	// 关闭socket
	defer listener.Close()
	// 启动监听广播的goroutine
	go this.ListenMsg()

	for{
		// 获取tcp握手成功的连接
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:",err)
		}

		// handler
		go this.Handler(conn)
	}

}