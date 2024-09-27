package ziface

/*
IRequest 接口：
实际上是把客户端请求的连接信息，和请求的数据包装到了一个 Request 中
*/

type IRequest interface {
	// GetConnection 得到当前连接
	GetConnection() IConnection

	// GetData 得到请求的消息数据
	GetData() []byte

	// GetMsgID 得到请求的消息的 ID
	GetMsgID() uint32
}
