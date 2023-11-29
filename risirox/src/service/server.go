package service

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net"
	"risirox/risirox/src/conf"
	iserver "risirox/risirox/src/server"
)

//定义全局的连接池管理器
var connManager iserver.IConnPoolManager

type Server struct {
	//服务器名称
	name string
	//服务器绑定的ip版本
	ipVersion string
	//ip
	ip string
	//端口
	port int
	//读写分离的写业务线程池
	workPool *WorkPool
	//消息的handler
	handler iserver.IMessageHandlerManager
	//连接启动前注册的方法
	preConnStartFunc func(conn iserver.IConnection)
	//连接销毁前注册的方法
	preConnDestroyFunc func(conn iserver.IConnection)
}

func (server *Server) Start() {
	fmt.Printf("server:[%s] ready to start\n", server.name)
	addr, err := net.ResolveTCPAddr(server.ipVersion, fmt.Sprintf("%s:%d", server.ip, server.port))
	if err != nil {
		log.Errorf("resolve tcp addr err:%s", err)
	}
	tcpListener, err := net.ListenTCP(server.ipVersion, addr)
	if err != nil {
		log.Errorf("listen ip:%s, port:%d, err:%s", server.ip, server.port, err)
	}
	var idx uint32 = 0
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Errorf("accept connection err:%s", err)
			continue
		}
		if connManager.GetPoolSize() >= conf.GlobalConfigObj.MaxConn {
			log.Info("Conn is refused!!")
			conn.Close()
			continue
		}
		if err != nil {
			log.Fatal(err)
		}
		connection := NewConnection(idx+1, conn, server)
		connection.Start()
		idx = idx + 1
	}
}

func (server *Server) Stop() {
	connManager.ClearConn()
	server.workPool.Close()
}

func (server *Server) Run() {
	server.Start()
}

func (server *Server) GetConnManager() iserver.IConnPoolManager {
	return connManager
}

func (server *Server) AddMsgHandler(msgType uint32, hanlder iserver.IMessageHandler) {
	server.handler.BindMsg2Handler(msgType, hanlder)
}

func (server *Server) GetMsgHandler() iserver.IMessageHandlerManager {
	return server.handler
}

func (server *Server) RegisterPreConnStartFunc(preConnStartFunc func(connection iserver.IConnection)) {
	server.preConnStartFunc = preConnStartFunc
}

func (server *Server) RegisterPreConnDestroyFunc(preConnDestroyFunc func(connection iserver.IConnection)) {
	server.preConnDestroyFunc = preConnDestroyFunc
}

func (server *Server) PreConnStart(conn iserver.IConnection) {
	if server.preConnStartFunc != nil {
		server.preConnStartFunc(conn)
	}
}

func (server *Server) PreConnDestroy(conn iserver.IConnection) {
	if server.preConnDestroyFunc != nil {
		server.preConnDestroyFunc(conn)
	}
}

func NewServer() iserver.IServer {
	connManager = NewConnPoolManager()
	return &Server{
		ip:        conf.GlobalConfigObj.Host,
		port:      conf.GlobalConfigObj.Port,
		ipVersion: "tcp4",
		name:      "RisiroX",
		handler:   NewConnHandler(),
	}
}
