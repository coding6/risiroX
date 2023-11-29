package iserver

type IMessage interface {
	// GetMsgId 获取消息Id
	GetMsgId() uint32

	// GetMsgLen 获取消息体长度
	GetMsgLen() uint32

	// GetData 获取消息内容
	GetData() []byte

	// SetMsgId 设置消息Id
	SetMsgId(id uint32)

	// SetMsgLen 设置消息长度
	SetMsgLen(msgLen uint32)

	// SetData 设置消息内容
	SetData(data []byte)

	// GetMsgType 设置消息类型
	GetMsgType() uint32
}
