package iserver

type IRequest interface {
	// GetConnection 得到连接对象
	GetConnection() IConnection

	// GetData 拿到当前请求的数据包
	GetData() []byte

	// GetMsgId 拿到当前请求消息的id
	GetMsgId() uint32

	// GetMsgType 获取当前消息类型
	GetMsgType() uint32
}
