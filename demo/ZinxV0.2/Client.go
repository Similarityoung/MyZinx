package main

import (
	"fmt"
	"net"
	"time"
)

/*
	模拟客户端
*/

func main() {
	fmt.Println("Client start...")

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Client start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("Hello Zinx V0.2..."))
		if err != nil {
			fmt.Println("write conn err ", err)
			return
		}

		bytes := make([]byte, 512)

		read, err := conn.Read(bytes)
		if err != nil {
			fmt.Println("read buf err ", err)
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n\n", bytes, read)
		time.Sleep(2 * time.Second)
	}
}
