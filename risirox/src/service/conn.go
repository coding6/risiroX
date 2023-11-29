package service

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	data2 "risirox/risirox/src/data"
	iserver "risirox/risirox/src/server"
)

type Connection struct {
	//链接id
	ConnId uint32
	//当前链接的socket句柄
	TcpConn *net.TCPConn
	//链接是否关闭标识
	isClose bool
	//告知当前链接关闭的消息管道
	ExitChan chan bool
	//消息传递的管道
	MsgData chan []byte
	//当前链接绑定的server
	Server iserver.IServer
	//当前链接绑定的message handler
	HandlerManager iserver.IMessageHandlerManager
}

// Start 开启链接
func (conn *Connection) Start() {
	tcpConn := conn.GetTcpConn()
	go conn.read(tcpConn)
	go conn.write(tcpConn)
	conn.GetServer().PreConnStart(conn)
}

func (conn *Connection) read(tcpConn *net.TCPConn) {
	for {
		msgHead := make([]byte, 12)
		//_, err := tcpConn.Read(data)
		_, err := io.ReadFull(tcpConn, msgHead)
		if err != nil {
			log.Errorf("read head err:%s", err)
			//read client message err, need to delete this connection in pool
			conn.Stop()
			break
		}
		//对消息进行解包
		packUtil := PackUtil{}
		msg, err := packUtil.UnPack(msgHead)
		var msgBody []byte
		if msg.GetMsgLen() > 0 {
			msgBody = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(tcpConn, msgBody)
			if err != nil {
				log.Errorf("read data err:%s", err)
				break
			}
		}
		msg.SetData(msgBody)
		request := NewRequest(conn, msg)
		//任务提交线程池执行message具体的逻辑
		conn.submitMsg2Pool(request)
	}
}

func (conn *Connection) write(tcpConn *net.TCPConn) {
	for {
		select {
		case msg := <-conn.MsgData:
			if _, err := tcpConn.Write(msg); err != nil {
				log.Errorf("write data err:%s", err)
				return
			}
		case <-conn.ExitChan:
			log.Infof("connection close singal is recved")
			return
		}
	}
}

func (conn *Connection) submitMsg2Pool(request iserver.IRequest) {
	task := &data2.Task{
		Handler: conn.HandlerManager.DoHandler,
		Param:   request,
	}
	WorkPoolObj.Submit(task)
}

func (conn *Connection) SendMsg2Client(msg iserver.IMessage) error {
	if conn.isClose {
		return errors.New("conn is closed")
	}
	packUtil := NewPackUtil()
	byteData, err := packUtil.Pack(msg)
	if err != nil {
		log.Errorf("pack message msgId:%s", msg.GetMsgId())
		return err
	}
	conn.MsgData <- byteData
	return nil
}

// Stop 关闭链接
func (conn *Connection) Stop() {
	if conn.isClose == true {
		return
	}
	conn.GetServer().PreConnDestroy(conn)
	conn.TcpConn.Close()
	conn.ExitChan <- true
	conn.GetServer().GetConnManager().DelConn(conn.GetConnId())
	close(conn.ExitChan)
	close(conn.MsgData)
}

// GetConnId 获取链接唯一标识
func (conn *Connection) GetConnId() uint32 {
	return conn.ConnId
}

func (conn *Connection) GetServer() iserver.IServer {
	return conn.Server
}

// GetTcpConn 获取当前链接的socket句柄
func (conn *Connection) GetTcpConn() *net.TCPConn {
	return conn.TcpConn
}

func NewConnection(connId uint32, tcpConn *net.TCPConn, server iserver.IServer) *Connection {
	conn := &Connection{
		ConnId:         connId,
		TcpConn:        tcpConn,
		Server:         server,
		isClose:        false,
		ExitChan:       make(chan bool),
		MsgData:        make(chan []byte),
		HandlerManager: server.GetMsgHandler(),
	}
	server.GetConnManager().AddConn(conn)
	return conn
}
