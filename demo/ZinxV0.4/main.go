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

	fmt.Printf("%s\n", request.GetData())
	err := request.GetConnection().SendMsg(1, []byte("PreHandle\n"))

	if err != nil {
		return
	}
}

func (r *Router) Handle(request ziface.IRequest) {
	err := request.GetConnection().SendMsg(1, []byte("Handle\n"))
	if err != nil {
		return
	}
}

func (r *Router) PostHandle(request ziface.IRequest) {
	err := request.GetConnection().SendMsg(1, []byte("PostHandle\n"))
	if err != nil {
		return
	}
}

func main() {

	s := znet.NewServer()

	// 注册路由
	s.AddRouter(&Router{})

	s.Serve()

}
