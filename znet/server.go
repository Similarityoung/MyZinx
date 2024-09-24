package znet

import (
	"fmt"
	"io"
	"net"
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
}

// Start  server
func (s *Server) Start() {
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

		// 阻塞的等待客户端链接。处理客户端链接业务（读写）
		for {
			tcp, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 客户端已经与客户端建立链接，做一些业务，做一个最基本的最大512字节长度的回写
			go func() {
				for {
					bytes := make([]byte, 512)
					read, err := tcp.Read(bytes)
					if err != nil && err != io.EOF {
						fmt.Println("read err", err)
						continue
					} else if err == io.EOF {
						return
					}

					fmt.Printf("recv from client, data: %s\n", bytes[:read])

					if _, err := tcp.Write(bytes[:read]); err != nil {
						fmt.Println("write err", err)
						continue
					}
				}
			}()
		}
	}()
}

// Stop server
func (s *Server) Stop() {

	// TODO
	// Do some server stop work, such as releasing the resource, etc.

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

// NewServer create a server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}

	return s
}
