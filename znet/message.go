package znet

type Message struct {
	ID   uint32
	Len  uint32
	Data []byte
}

func NewMessage(data []byte, ID uint32) *Message {
	return &Message{Data: data, ID: ID, Len: uint32(len(data))}
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) GetMsgID() uint32 {
	return m.ID
}

func (m *Message) GetMsgLen() uint32 {
	return m.Len
}

func (m *Message) SetData(bytes []byte) {
	m.Data = bytes
}

func (m *Message) SetMsgID(u uint32) {
	m.ID = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.Len = u
}
