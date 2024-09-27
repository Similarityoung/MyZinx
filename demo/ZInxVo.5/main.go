package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

type Router struct {
	router znet.BaseRouter
}

func (r *Router) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	fmt.Printf("MsgID = %d ,data = %s\n", request.GetMsgID(), request.GetData())

	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping...\n"))

	if err != nil {
		fmt.Println("call back ping...ping...ping error:", err)
		return
	}
}

func (r *Router) Handle(request ziface.IRequest) {

}

func (r *Router) PostHandle(request ziface.IRequest) {

}

func main() {

	s := znet.NewServer()

	// 注册路由
	s.AddRouter(&Router{})

	s.Serve()

}
