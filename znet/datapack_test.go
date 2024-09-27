package znet

import (
	"io"
	"net"
	"sync"
	"testing"
)

// 单元测试
func TestDataPack(t *testing.T) {
	// 模拟服务器
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatalf("net.Listen err: %v", err)
	}
	defer listener.Close() // 确保监听器关闭

	var wg sync.WaitGroup

	// 模拟服务器接收客户端请求
	go func() {
		defer wg.Done()
		conn, err := listener.Accept()
		if err != nil {
			t.Errorf("listener.Accept err: %v", err)
			return
		}
		defer conn.Close()

		handleConnection(conn, t)
	}()

	// 模拟客户端
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatalf("net.Dial err: %v", err)
	}
	defer conn.Close() // 确保连接关闭

	// 封包
	dp := NewDataPack()
	msg1 := &Message{
		ID:   1,
		Data: []byte("hello"),
		Len:  uint32(len([]byte("hello"))),
	}

	data1, err := dp.Pack(msg1)
	if err != nil {
		t.Fatalf("dp.Pack err: %v", err)
	}

	msg2 := &Message{
		ID:   1,
		Data: []byte("helloZinx"),
		Len:  uint32(len([]byte("helloZinx"))),
	}

	data2, err := dp.Pack(msg2)
	if err != nil {
		t.Fatalf("dp.Pack err2: %v", err)
	}

	// 合并数据，模拟粘包
	bytes := append(data1, data2...)
	_, err = conn.Write(bytes)
	if err != nil {
		t.Fatalf("conn.Write err: %v", err)
	}

	// 等待服务器协程处理完成
	wg.Wait()
}

func handleConnection(conn net.Conn, t *testing.T) {
	dp := NewDataPack()

	for {
		// 读取客户端的消息头
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, headData); err != nil {
			if err != io.EOF {
				t.Errorf("conn.Read headData err: %v", err)
			}
			return // 从循环中退出
		}

		// 拆包
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			t.Errorf("dp.Unpack err: %v", err)
			return
		}

		msg, ok := msgHead.(*Message)
		if !ok {
			t.Errorf("message type assertion failed")
			return
		}

		if msg.GetMsgLen() > 0 {
			// 读取客户端的消息体
			msgData := make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msgData); err != nil {
				t.Errorf("conn.Read msgData err: %v", err)
				return
			}

			// 设置消息体
			msg.SetData(msgData)

			// 打印消息
			t.Logf("msg.GetMsgID() = %v, msg.GetMsgLen() = %v, msg.GetData() = %v",
				msg.GetMsgID(), msg.GetMsgLen(), string(msg.GetData()))
		}
	}
}
