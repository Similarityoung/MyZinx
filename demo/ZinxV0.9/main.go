package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping...Handle\n"))
	if err != nil {
		fmt.Println("call back ping...ping...ping error:", err)
		return
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
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

func main() {

	s := znet.NewServer()

	// 注册路由
	s.AddRouter(0, &PingRouter{})

	s.AddRouter(1, &HelloZinxRouter{})

	s.Serve()

}
