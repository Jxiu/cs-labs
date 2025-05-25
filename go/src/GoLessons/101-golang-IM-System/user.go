package main

import (
	"net"
	// "fmt"
	"strings"
)
type User struct{
	Name string
	Addr string
	// 当前用户的channel
	C chan string 
	// 跟客户端的连击
	conn net.Conn
	// 所属Server
	server *Server
}
// 用户上线业务
func (this *User) Online(){

	// 加入 onlineMap
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播用户上线消息
	this.server.BroadCast(this, "Online!!!")
}
// 用户下线业务
func (this *User) Offline(){
	// 加入 onlineMap
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播用户上线消息
	this.server.BroadCast(this, "Offline---")
}

// 给当前用户对应的客户端发送消息
func (this *User) SendMsg(msg string){
	this.conn.Write([]byte(msg))
}

// 用户处理消息
func (this *User) DoMsg(msg string){
	msg = strings.TrimSpace(msg) // 去除空白
	if msg == "who"{
		this.server.mapLock.Lock()
		for _,user := range this.server.OnlineMap{
			onlineMsg := "["+user.Addr+"]" + user.Name + " : " + "is online \n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	}else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := strings.Split(msg,"|")[1]
		// 判断name是否存在
		_,ok := this.server.OnlineMap[newName]
		if ok {
			this.SendMsg("current user name exist\n")
		}else{
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap,this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("Your New Uername is: " + newName + "\n")
		}
	}else if len(msg)>3 && msg[:3] == "to|" {
		// 消息格式：to|张三|消息内容
		// 获取对方用户名
		remoteName := strings.Split(msg,"|")[1]
		if remoteName == "" {
			this.SendMsg("Message format not support, send to others format like: \"to|jack|hello!\" \n")
			return
		}
		// 找到用户发送消息
		remoteUser,ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("remote user not exist!!")
			return
		}
		content := strings.Split(msg,"|")[2]
		if content == ""{
			this.SendMsg("content is empty, please retry")
			return
		}
			remoteUser.SendMsg( "[whisper]:" + this.Name + ":" + content + "\n")
		
	}else{
		this.server.BroadCast(this, msg)
	}
}
// 监听当前User channel的方法，一旦有消息就直接发给客户端
// 有其他用户写到当前用户channel 中，转发给用户
func (this *User) ListenMsg(){
	for{
		msg := <- this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User {
		Name : userAddr,
		Addr : userAddr,
		C : make(chan string),
		conn : conn,
		server: server,
	}
	// 启动监听管道
	go user.ListenMsg()
	return user
}

