package xinterface

// IServer define server interface
type IServer interface {

	// start server
	Start()

	// stop server
	Stop()

	// start serve
	Serve()

	// add router
	AddRouter(msgID uint32, router IRouter)

	GetConnMgr() IConnMgr

	SetOnConnStart(func(IConnection))

	SetOnConnStop(func(IConnection))

	CallOnConnStart(conn IConnection)

	CallOnConnStop(conn IConnection)
}
