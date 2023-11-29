package iserver

type IServer interface {
	// Run 运行服务器，对外暴露
	Run()

	// Start 启动服务器
	Start()

	// Stop 暂停服务器
	Stop()

	// GetConnManager 获取Server的链接管理器
	GetConnManager() IConnPoolManager

	// AddMsgHandler 给指定message绑定handler处理器
	AddMsgHandler(msgType uint32, hanlder IMessageHandler)

	// GetMsgHandler 获取消息处理管理器对象
	GetMsgHandler() IMessageHandlerManager

	RegisterPreConnStartFunc(preConnStartFunc func(connection IConnection))

	RegisterPreConnDestroyFunc(preConnDestroyFunc func(connection IConnection))

	PreConnStart(conn IConnection)

	PreConnDestroy(conn IConnection)
}
