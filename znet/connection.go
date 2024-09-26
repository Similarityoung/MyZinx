package znet

import (
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// Connection 接口实现
type Connection struct {
	// 当前链接的 socket TCP 套接字
	Conn *net.TCPConn

	// 链接的 ID
	ConnID uint32

	// 当前链接的状态
	isClosed bool

	// 告知当前链接已经退出/停止的 channel
	ExitChan chan bool

	// 该链接处理的方法 Router
	Router ziface.IRouter
}

// NewConnection 初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

// Start connection
func (connection *Connection) Start() {
	fmt.Println("Conn Start()...ConnID = ", connection.ConnID)

	// 启动当前链接的读数据业务
	go connection.StartReader()
}

// Stop connection
func (connection *Connection) Stop() {
	fmt.Println("Conn Stop()...ConnID = ", connection.ConnID)

	if connection.isClosed {
		return
	}
	connection.isClosed = true

	// 关闭 socket 链接
	err := connection.Conn.Close()
	if err != nil {
		return
	}

	// 关闭当前链接全部管道
	close(connection.ExitChan)
}

// GetConnection 获取当前链接绑定的socket conn
func (connection *Connection) GetConnection() *net.TCPConn {
	return connection.Conn
}

// GetConnID 获取当前链接模块的链接ID
func (connection *Connection) GetConnID() uint32 {
	return connection.ConnID
}

// RemoteAddr 获取远程客户端的 TCP 状态 IP port
func (connection *Connection) RemoteAddr() net.Addr {
	return connection.Conn.RemoteAddr()
}

// SendMsg 发送数据，将数据发送给远程的客户端
func (connection *Connection) SendMsg(msgId uint32, data []byte) error {
	_, err := connection.Conn.Write(data)
	if err != nil {
		fmt.Println("SendMsg error", err)
		return err
	}
	return nil
}

// StartReader 启动读协程
func (connection *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running...")
	defer fmt.Println("connID = ", connection.ConnID, "Reader is exit, remote addr is ", connection.RemoteAddr().String())
	defer connection.Stop()

	for {
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		// 必须读到从客户端发来的数据才可以进行下一步处理
		_, err := connection.Conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("recv buf err", err)
			continue
		} else if err == io.EOF {
			return
		}

		//// 调用当前链接所绑定的 HandleAPI
		//if err := connection.handleAPI(connection.Conn, buf, read); err != nil {
		//	fmt.Println("ConnID", connection.ConnID, "handle is error", "err is ", err)
		//	break
		//}

		// 得到当前数据的 request 请求
		// 这个 buf 是服务器读出来的数据，由客户端发送过来的，现在要把它交给路由，让路由处理
		request := &Request{
			conn: connection,
			data: buf,
		}

		// 执行注册路由的 Handle 方法
		go func() {
			connection.Router.PreHandle(request)
			connection.Router.Handle(request)
			connection.Router.PostHandle(request)
		}()
	}

}
