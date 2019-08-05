package xnet

import (
	"fmt"
	"goserv/util"
	"goserv/xinterface"
)

// MsgHandle ..
type MsgHandle struct {
	Apis map[uint32]xinterface.IRouter

	WorkerPoolSize uint32

	TaskQueue []chan xinterface.IRequest
}

// NewMsgHandle ..
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]xinterface.IRouter),
		WorkerPoolSize: util.GlobalObject.WorkPoolSize,
		TaskQueue:      make([]chan xinterface.IRequest, util.GlobalObject.WorkPoolSize),
	}
}

// DoMsgHandle ..
func (mh *MsgHandle) DoMsgHandle(request xinterface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("Not found msg handle ", request.GetMsgID())
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

// startOneWorker ..
func (mh *MsgHandle) startOneWorker(workerID int, taskQ chan xinterface.IRequest) {
	fmt.Println("Start a new worker:", workerID)

	for {
		select {
		case req := <-taskQ:
			mh.DoMsgHandle(req)
		}
	}
}

// StartWorkerPool ...
func (mh *MsgHandle) StartWorkerPool() {

	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan xinterface.IRequest, util.GlobalObject.MaxTaskLen)
		go mh.startOneWorker(i, mh.TaskQueue[i])
	}
}

// SendMsgToTaskQueue ...
func (mh *MsgHandle) SendMsgToTaskQueue(request xinterface.IRequest) {

	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize

	fmt.Println("Add connID=", request.GetConnection().GetConnID(), "requestID=",
		request.GetMsgID(), "to workerID=", workerID)

	mh.TaskQueue[workerID] <- request
}
