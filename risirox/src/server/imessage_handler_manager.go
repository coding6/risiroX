package iserver

// IMessageHandlerManager 连接处理管理器
/**
description: 消息需要绑定处理器，比如消息类型是普通消息，那么用户可以写一个handler做普通消息的处理
，并将普通消息调用BindMsg2Handler绑定指定的handler
*/
type IMessageHandlerManager interface {
	// BindMsg2Handler 给消息绑定处理器
	BindMsg2Handler(msgType uint32, handler IMessageHandler)

	// DoHandler 根据message执行handler操作
	DoHandler(request IRequest)
}
