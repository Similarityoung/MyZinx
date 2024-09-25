package ziface

import "net"

// IConnection 定义链接模块的抽象层
type IConnection interface {
	// Start connection
	Start()

	// Stop connection
	Stop()

	// GetConnection 获取当前链接绑定的socket conn
	GetConnection() *net.TCPConn

	// GetConnID 获取当前链接模块的链接ID
	GetConnID() uint32 // 无符号整型，32位 4 字节

	// RemoteAddr 获取远程客户端的 TCP 状态 IP port
	RemoteAddr() net.Addr

	// SendMsg 发送数据，将数据发送给远程的客户端
	SendMsg(msgId uint32, data []byte) error

	//// SetProperty 设置链接属性
	//SetProperty(key string, value interface{})
}

// HandleFunc 定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn, []byte, int) error
