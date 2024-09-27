package ziface

// IMessage interface
type IMessage interface {

	// GetData 获取消息的信息
	GetData() []byte
	// GetMsgID 获取消息的ID
	GetMsgID() uint32
	// GetMsgLen 获取消息的长度
	GetMsgLen() uint32

	// SetData 设置消息的信息
	SetData([]byte)
	// SetMsgID 设置消息的ID
	SetMsgID(uint32)
	// SetMsgLen 设置消息的长度
	SetMsgLen(uint32)
}
