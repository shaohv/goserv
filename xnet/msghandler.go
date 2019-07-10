package xnet

import (
	"fmt"
	"goserv/xinterface"
)

// MsgHandle ..
type MsgHandle struct {
	Apis map[uint32]xinterface.IRouter
}

// NewMsgHandle ..
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]xinterface.IRouter),
	}
}

// DoMsgHandle ..
func (mh *MsgHandle) DoMsgHandle(request xinterface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("Not found msg %d handle ",
			request.GetMsgID())
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter ..
func (mh *MsgHandle) AddRouter(msgID uint32, route xinterface.IRouter) {
	if _, ok := mh.Apis[msgID]; ok {
		fmt.Printf("msgID %d already registed!", msgID)
		return
	}

	mh.Apis[msgID] = route
	fmt.Printf("add route for msgID %d", msgID)
}
