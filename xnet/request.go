package xnet

import (
	"goserv/xinterface"
)

// Request represent a client req
type Request struct {
	conn xinterface.IConnection
	//data []byte
	msg xinterface.IMessage
}

// GetConnection return req's connn
func (r *Request) GetConnection() xinterface.IConnection {
	return r.conn
}

// GetData return req's data
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// GetMsgID ..
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
