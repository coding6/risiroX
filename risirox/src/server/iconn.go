package iserver

import "net"

/**
关系的抽象类
*/
type IConnection interface {
	// Start 启动链接
	Start()
	// Stop 关闭链接
	Stop()
	// GetConnId 获取链接唯一标识
	GetConnId() uint32
	// GetTcpConn 获取当前链接的socket句柄
	GetTcpConn() *net.TCPConn

	GetServer() IServer

	SendMsg2Client(msg IMessage) error
}
