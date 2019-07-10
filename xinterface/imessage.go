package xinterface

// IMessage ..
type IMessage interface {
	GetDataLen() uint32 // GetDataLen ..

	GetMsgId() uint32 // GetMsgId ..

	GetData() []byte // GetData ..

	SetMsgId(uint32) // SetMsgId ..

	SetData([]byte) // SetData ..

	SetDataLen(uint32) // SetDataLen ..
}
