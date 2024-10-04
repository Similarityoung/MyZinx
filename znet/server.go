package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// Server iServer interface define a server struct
type Server struct {
	// Server name
	Name string

	// Server IP version, ipv4 or ipv6
	IPVersion string

	// Server IP
	IP string

	// Server port
	Port int

	// MsgHandler
	MsgHandler ziface.IMessageHandler

	// ConnManager
	ConnManager ziface.IConnManager

	// OnConnStart callback
	OnConnStart func(conn ziface.IConnection)

	// OnConnStop callback
	OnConnStop func(conn ziface.IConnection)
}

func (s *Server) GetConnManager() ziface.IConnManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

func (s *Server) SetOnConnStop(hookFunc func(conn ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}

// CallBackToClient 当前客户端链接的所绑定的 API,后续由框架使用者自行开发
//func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
//fmt.Println("[Conn Handle] CallBackToClient ...")
//if _, err := conn.Write(data[:cnt]); err != nil {
//	fmt.Println("write back buf err", err)
//	return errors.New("CallBackToClient error")
//}
//return nil
//}

// Start  server
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, listenner at IP: %s, Port %d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)
	fmt.Printf("[Start] Server Listenner at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	// 开启一个 go 承载服务,因为监听只是服务的一个业务，还有其他的业务需要实现，所以监听开个协程来进行
	go func() {
		// 获取 一个 tcp 的 addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error", err)
			return
		}
		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server succ, ", s.Name, " succ, Listening...")
		var cid uint32
		cid = 0

		// 阻塞的等待客户端链接。处理客户端链接业务（读写）
		for {
			tcp, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			if s.GetConnManager().Len() >= utils.GlobalObject.MaxConn {
				// TODO 给客户端响应一个超出最大连接的错误包
				fmt.Println("Too many connections, MaxConn = ", utils.GlobalObject.MaxConn)
				err := tcp.Close()
				if err != nil {
					fmt.Println("Close err", err)
					return
				}
				continue
			}

			// 将处理连接的业务方法和 conn 进行绑定，得到我们的连接模块
			connection := NewConnection(s, tcp, cid, s.MsgHandler)
			cid++
			connection.Start()
		}
	}()
}

// Stop server
func (s *Server) Stop() {

	fmt.Println("[STOP] Zinx server name ", s.Name)
	s.GetConnManager().ClearConn()

}

// Serve Run server
func (s *Server) Serve() {
	// Start server
	s.Start()

	// TODO
	// Do some server initialization work, such as reading configuration, initializing the database, etc.

	// 阻塞
	select {}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("Add Router succ!")
}

// NewServer create a server
func NewServer() ziface.IServer {
	s := &Server{
		Name:        utils.GlobalObject.Name,
		IPVersion:   "tcp4",
		IP:          utils.GlobalObject.Host,
		Port:        utils.GlobalObject.TcpPort,
		MsgHandler:  NewMessageHandler(),
		ConnManager: NewConnManager(),
	}

	// 初始化 WorkerPool
	s.MsgHandler.StartWorkerPool()

	return s
}
