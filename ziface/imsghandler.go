package ziface

type IMessageHandler interface {
	// DoMsgHandler 调度/执行对应的 msgHandler 消息处理方法
	DoMsgHandler(request IRequest)

	// AddRouter 为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)

	// StartWorkerPool 启动 Worker 工作池
	StartWorkerPool()

	// SendMsgToTaskQueue 将消息交给 TaskQueue，由 Worker 进行处理
	SendMsgToTaskQueue(request IRequest)
}
