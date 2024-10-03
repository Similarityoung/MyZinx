package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	router znet.BaseRouter
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call msgHandler PreHandle")
	fmt.Printf("MsgID = %d ,data = %s\n", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping...\n"))

	if err != nil {
		fmt.Println("call back ping...ping...ping error:", err)
		return
	}
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping...Handle\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error:", err)
		return
	}
}

func (r *PingRouter) PostHandle(request ziface.IRequest) {

}

type HelloZinxRouter struct {
	router znet.BaseRouter
}

func (h *HelloZinxRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call msgHandler PreHandle")
	fmt.Printf("MsgID = %d ,data = %s\n", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(1, []byte("Hello...hello...hello...\n"))

	if err != nil {
		fmt.Println("call back ping...ping...ping error:", err)
		return
	}
}

func (h *HelloZinxRouter) Handle(request ziface.IRequest) {

}

func (h *HelloZinxRouter) PostHandle(request ziface.IRequest) {

}

func main() {

	s := znet.NewServer()

	// 注册路由
	s.AddRouter(0, &PingRouter{})

	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()

}
