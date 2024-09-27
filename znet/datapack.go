package znet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"zinx/utils"
	"zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	// DataLen uint32(4 byte) + ID uint32(4 byte)
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	// 创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 将dataLen写进dataBuff中
	err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen())
	if err != nil {
		fmt.Println("Pack binary.Write dataLen error:", err)
		return nil, err
	}

	// 将msgID写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgID())
	if err != nil {
		fmt.Println("Pack binary.Write msgID error:", err)
		return nil, err
	}

	// 将data写进dataBuff中
	err = binary.Write(dataBuff, binary.LittleEndian, msg.GetData())
	if err != nil {
		fmt.Println("Pack binary.Write data error:", err)
		return nil, err
	}

	return dataBuff.Bytes(), nil

}

func (d *DataPack) Unpack(byte []byte) (ziface.IMessage, error) {
	// 创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(byte)

	// 初始化一个IMessage
	msg := &Message{}

	// 读dataLen
	err := binary.Read(dataBuff, binary.LittleEndian, &msg.Len)
	if err != nil {
		fmt.Println("Unpack binary.Read dataLen error:", err)
		return nil, err
	}

	// 读msgID
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.ID)
	if err != nil {
		fmt.Println("Unpack binary.Read msgID error:", err)
		return nil, err
	}

	// 判断是否超出了我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.Len > utils.GlobalObject.MaxPackageSize {
		return nil, fmt.Errorf("too large msg data received")
	}

	return msg, nil
}
