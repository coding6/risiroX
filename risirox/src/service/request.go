package service

import iserver "risirox/risirox/src/server"

type Request struct {
	Conn iserver.IConnection

	Message iserver.IMessage
}

// GetConnection 得到当前绑定的链接数据
func (request *Request) GetConnection() iserver.IConnection {
	return request.Conn
}

// GetData 拿到当前请求的数据包
func (request *Request) GetData() []byte {
	return request.Message.GetData()
}

// GetMsgId 拿到当前请求消息的id
func (request *Request) GetMsgId() uint32 {
	return request.Message.GetMsgId()
}

func (request *Request) GetMsgType() uint32 {
	return request.Message.GetMsgType()
}

func NewRequest(conn iserver.IConnection, message iserver.IMessage) iserver.IRequest {
	return &Request{
		Conn:    conn,
		Message: message,
	}
}
