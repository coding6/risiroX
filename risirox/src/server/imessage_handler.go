package iserver

type IMessageHandler interface {
	// Handler 处理msg具体的业务逻辑
	Handler(request IRequest)
}
