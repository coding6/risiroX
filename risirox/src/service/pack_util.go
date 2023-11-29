package service

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"risirox/risirox/src/conf"
	iserver "risirox/risirox/src/server"
)

type PackUtil struct {
}

func (packUtil *PackUtil) Pack(msg iserver.IMessage) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	//将dataLen写入缓冲区
	err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		fmt.Println("write datalen err")
		return nil, err
	}
	//将msgType写入buffer
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgType()); err != nil {
		fmt.Println("write msg type err")
		return nil, err
	}
	//将msgId写入buffer
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetMsgId()); err != nil {
		fmt.Println("write msgId err")
		return nil, err
	}
	//将data写入buffer
	if err := binary.Write(buffer, binary.LittleEndian, msg.GetData()); err != nil {
		fmt.Println("write data err")
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (packUtil *PackUtil) UnPack(data []byte) (iserver.IMessage, error) {
	buffer := bytes.NewReader(data)
	message := &Message{}
	if err := binary.Read(buffer, binary.LittleEndian, &message.DataLen); err != nil {
		fmt.Println("read msgLen err:", err, "len:", message.GetMsgLen())
		return nil, err
	}

	if err := binary.Read(buffer, binary.LittleEndian, &message.MsgType); err != nil {
		fmt.Println("read MsgType err:", err)
		return nil, err
	}

	if err := binary.Read(buffer, binary.LittleEndian, &message.Id); err != nil {
		fmt.Println("read msgId err:", err)
		return nil, err
	}

	if message.GetMsgLen() > conf.GlobalConfigObj.MaxPackageSize {
		return nil, errors.New("too large data in this conn")
	}
	return message, nil
}

func NewPackUtil() *PackUtil {
	return &PackUtil{}
}
