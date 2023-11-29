package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"risirox/risirox/src/data"
	"risirox/risirox/src/service"
	"time"
)

func main() {
	time.Sleep(1 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Errorf("err: %s", err)
		return
	}

	for {
		fmt.Println("send msg ...")
		util := service.NewPackUtil()
		pack, _ := util.Pack(service.NewMsg(123456, []byte("wo shi ni baba hhhhhhhhhhh"), data.NormalMessage))
		_, err2 := conn.Write(pack)

		msgHead := make([]byte, 12)
		//_, err := tcpConn.Read(data)
		_, err := io.ReadFull(conn, msgHead)
		if err != nil {
			log.Errorf("read head err:%s", err)
			break
		}
		//对消息进行解包
		packUtil := service.PackUtil{}
		msg, err := packUtil.UnPack(msgHead)
		var msgBody []byte
		if msg.GetMsgLen() > 0 {
			msgBody = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(conn, msgBody)
			if err != nil {
				log.Errorf("read data err:%s", err)
				break
			}
			fmt.Println("msg from client:", string(msgBody), ",msgId:", msg.GetMsgId())
		}
		if err2 != nil {
			log.Errorf("err:%s", err2)
			break
		}
		time.Sleep(5 * time.Second)
	}
}
