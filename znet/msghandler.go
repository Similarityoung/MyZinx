package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

type MessageHandler struct {
	// 存放每个 MsgID 所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{Apis: make(map[uint32]ziface.IRouter)}
}

func (m *MessageHandler) DoMsgHandler(request ziface.IRequest) {
	// 从 Request 中找到 msgID
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgID = ", request.GetMsgID(), " is NOT FOUND! Need Register!")
		return
	}

	// 根据 msgID 调度对应的 MsgHandler 业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (m *MessageHandler) AddRouter(msgID uint32, router ziface.IRouter) {
	// 1. 判断当前 msg 绑定的 API 处理方法是否已经存在
	if _, ok := m.Apis[msgID]; ok {
		panic("repeat api, msgID = " + strconv.Itoa(int(msgID)))
	}

	// 2. 添加 msg 与 API 的绑定关系
	m.Apis[msgID] = router
	fmt.Println("Add api msgID = ", msgID)
}
