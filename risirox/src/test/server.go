package main

import (
	"fmt"
	"risirox/risirox/src/data"
	"risirox/risirox/src/logo"
	iserver "risirox/risirox/src/server"
	"risirox/risirox/src/service"
)

var s iserver.IServer

type ServerMessageHandler struct {
	service.BaseMessageHandler
}

func (handler *ServerMessageHandler) Handler(request iserver.IRequest) {
	//将收到的消息广播到其他客户端
	request.GetConnection()
	connMap := s.GetConnManager().GetAllConn()
	for _, conn := range connMap {
		msg := service.NewMsg(request.GetMsgId(), request.GetData(), request.GetMsgType())
		conn.SendMsg2Client(msg)
	}
}

func PreConnStart(conn iserver.IConnection) {
	fmt.Println("连接启动了")
}

func PreDestroy(conn iserver.IConnection) {
	fmt.Println("连接销毁了")
}

func main() {
	logo.PrintLogo()
	s = service.NewServer()
	s.AddMsgHandler(data.NormalMessage, &ServerMessageHandler{})
	s.RegisterPreConnStartFunc(PreConnStart)
	s.RegisterPreConnDestroyFunc(PreDestroy)
	s.Run()
}
