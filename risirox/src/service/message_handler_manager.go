package service

import (
	"fmt"
	iserver "risirox/risirox/src/server"
)

type MessageHandlerManager struct {
	HandlerMap map[uint32]iserver.IMessageHandler
}

func NewConnHandler() *MessageHandlerManager {
	return &MessageHandlerManager{
		HandlerMap: make(map[uint32]iserver.IMessageHandler),
	}
}

func (connHandler *MessageHandlerManager) BindMsg2Handler(_type uint32, handler iserver.IMessageHandler) {
	if _, ok := connHandler.HandlerMap[_type]; ok {
		fmt.Println("msgType:", _type, "had bind a hookHandler, please do not bind again")
		return
	}
	connHandler.HandlerMap[_type] = handler
}

func (connHandler *MessageHandlerManager) DoHandler(request iserver.IRequest) {
	handler, ok := connHandler.HandlerMap[request.GetMsgType()]
	if !ok {
		return
	}
	handler.Handler(request)
}
