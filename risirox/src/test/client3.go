package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
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
		msgHead := make([]byte, 8)
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
			fmt.Println("msg from client:", string(msgBody))
		}
		time.Sleep(5 * time.Second)
	}
}
