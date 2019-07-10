package xnet

type Message struct {
	Id uint32

	DataLen uint32

	Data []byte
}

// NewMsgPkg ..
func NewMsgPkg(id uint32, data []byte) *Message {
	msg := &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}

	return msg
}

// GetMsgId ..
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

// GetDataLen ..
func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

// GetData ..
func (m *Message) GetData() []byte {
	return m.Data
}

// SetMsgId ..
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

// SetDataLen ..
func (m *Message) SetDataLen(l uint32) {
	m.DataLen = l
}

// SetData ..
func (m *Message) SetData(data []byte) {
	m.Data = data
}
