package xinterface

// IServer define server interface
type IServer interface {

	// start server
	Start()

	// stop server
	Stop()

	// start serve
	Serve()
}
