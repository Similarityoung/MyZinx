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

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionBegin is Called ...")
	if err := conn.SendMsg(202, []byte("DoConnection BEGIN")); err != nil {
		fmt.Println(err)
	}

	// 给当前链接设置一些属性
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "ZinxV1.0")
	conn.SetProperty("Home", "https://simi.host")
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("===> DoConnectionLost is Called ...")
	fmt.Println("Conn ID = ", conn.GetConnID(), " is LOST")

	// 获取连接属性
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}

	if home, err := conn.GetProperty("Home"); err == nil {
		fmt.Println("Conn Property Home = ", home)
	}
}

func main() {

	s := znet.NewServer()

	// 注册路由
	s.AddRouter(0, &PingRouter{})

	s.AddRouter(1, &HelloZinxRouter{})

	s.SetOnConnStart(DoConnectionBegin)

	s.SetOnConnStop(DoConnectionLost)

	s.Serve()

}
