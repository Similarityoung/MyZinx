package znet

import (
	"fmt"
	"strconv"
	"zinx/utils"
	"zinx/ziface"
)

type MessageHandler struct {
	// 存放每个 MsgID 所对应的处理方法
	Apis map[uint32]ziface.IRouter
	// 负责 Worker 取任务的消息队列
	TaskQueue []chan ziface.IRequest
	// 业务工作 Worker 池的 worker 数量
	WorkerPoolSize uint32
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
	}
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

// StartWorkerPool 启动 Worker 工作池
func (m *MessageHandler) StartWorkerPool() {
	// 根据 WorkerPoolSize 分别开启 Worker，每个 Worker 用一个 go 承载
	for i := 0; i < int(m.WorkerPoolSize); i++ {
		// 1. 给当前的 Worker 对应的 channel 消息队列开辟空间，第 0 个 Worker 就用第 0 个 channel
		m.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)
		// 2. 启动当前的 Worker，阻塞等待消息从对应的 channel 传递进来
		go m.StartOneWorker(i, m.TaskQueue[i])
	}
}

// StartOneWorker 启动一个 Worker 工作流程
func (m *MessageHandler) StartOneWorker(workerID int, taskQueue chan ziface.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")

	// 不断阻塞等待对应消息队列的消息
	for {
		select {
		// 如果有消息过来，出列的就是一个客户端的 Request，执行当前 Request 所绑定的业务
		case request := <-taskQueue:
			m.DoMsgHandler(request)
		}
	}
}

// SendMsgToTaskQueue 将消息交给 TaskQueue，由 Worker 进行处理
func (m *MessageHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	// 1. 将消息平均分配给不同的 Worker
	// 根据客户端建立的 ConnID 来进行分配
	workerID := request.GetConnection().GetConnID() % m.WorkerPoolSize
	fmt.Println("Add ConnID = ", request.GetConnection().GetConnID(), " request MsgID = ", request.GetMsgID(),
		" to WorkerID = ", workerID)
	// 2. 将消息发送给对应的 Worker 的 TaskQueue
	m.TaskQueue[workerID] <- request
}
