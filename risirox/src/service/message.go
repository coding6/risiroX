package service

type Message struct {
	//消息id
	Id uint32

	//消息长度
	DataLen uint32

	//消息内容
	Data []byte

	//消息类型
	MsgType uint32
}

func NewMsg(msgId uint32, data []byte, msgType uint32) *Message {
	return &Message{
		Id:      msgId,
		Data:    data,
		DataLen: uint32(len(data)),
		MsgType: msgType,
	}
}

func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

func (msg *Message) SetMsgId(id uint32) {
	msg.Id = id
}

func (msg *Message) GetMsgLen() uint32 {
	return msg.DataLen
}

func (msg *Message) SetMsgLen(msgLen uint32) {
	msg.DataLen = msgLen
}

func (msg *Message) GetData() []byte {
	return msg.Data
}

func (msg *Message) SetData(data []byte) {
	msg.Data = data
}

func (msg *Message) GetMsgType() uint32 {
	return msg.MsgType
}
