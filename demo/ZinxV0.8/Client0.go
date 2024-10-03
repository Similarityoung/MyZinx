package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

/*
	模拟客户端
*/

func main() {
	fmt.Println("Client start...")

	time.Sleep(1 * time.Second)

	// 连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("Client start err, exit!")
		return
	}

	for {
		// 写,封包消息
		dp := znet.NewDataPack()
		msg, err := dp.Pack(znet.NewMessage([]byte("ZinxV0.6 client Test0 Message"), 0))
		if err != nil {
			fmt.Println("write conn err ", err)
			break
		}

		_, err = conn.Write(msg)
		if err != nil {
			fmt.Println("write conn err ", err)
			break
		}

		bytes := make([]byte, dp.GetHeadLen())

		// 这是客户端的 conn，用来读取服务器返回的数据
		_, err = io.ReadFull(conn, bytes)
		if err != nil {
			fmt.Println("read head err ", err)
			break
		}
		// 二次读主体
		msgHead, err := dp.Unpack(bytes)
		if err != nil {
			fmt.Println("server unpack err:", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())

			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				break
			}

			fmt.Printf("==> Recv Msg: ID=%d, len=%d, data=%s\n", msg.ID, msg.Len, msg.Data)
		}

		time.Sleep(5 * time.Second)
	}
}
