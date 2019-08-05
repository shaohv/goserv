package xinterface

// IMsgHandle ..
type IMsgHandle interface {
	DoMsgHandle(request IRequest)

	AddRouter(msgID uint32, route IRouter)

	StartWorkerPool()

	SendMsgToTaskQueue(request IRequest)
}
