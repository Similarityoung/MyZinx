package znet

import "zinx/ziface"

type Request struct {
	// 已经和客户端建立好的连接
	// 这是服务器的连接
	conn ziface.IConnection

	// 客户端请求的数据
	message ziface.IMessage
}

func (r *Request) GetMsgID() uint32 {
	return r.message.GetMsgID()
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.message.GetData()
}
