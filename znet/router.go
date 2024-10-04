package znet

import (
	"fmt"
	"zinx/ziface"
)

// BaseRouter 实现 ziface.IRouter 接口时，先嵌入这个基类，然后根据需要对这个基类的方法进行重写
// 这样就不需要每次实现 IRouter 接口时，都要实现 PreHandle、Handle、PostHandle 这三个方法
// 这样的好处是，如果以后需要对 IRouter 接口增加方法，只需要在 BaseRouter 中增加方法即可
// 路由中 request 的 PreHandle、Handle、PostHandle 方法的实现，都是针对服务器的连接的，写操作是给客户端的，读操作是客户端给自己的
type BaseRouter struct{}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("无事发生 PrHandle")
	_ = request.GetConnection().SendMsg(1, []byte("无事发生 PrHandle\n"))
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
	fmt.Println("无事发生 Handle")
	_ = request.GetConnection().SendMsg(1, []byte("无事发生 Handle\n"))
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("无事发生 PostHandle")
	_ = request.GetConnection().SendMsg(1, []byte("无事发生 PostHandle\n"))
}
