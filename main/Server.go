package main

import (
	"fmt"
	"goserv/xinterface"
	"goserv/xnet"
)

// PingRouter is a applic router
type PingRouter struct {
	xnet.BaseRouter
}

// PreHandle ..
func (pingRouter *PingRouter) PreHandle(request xinterface.IRequest) {
	fmt.Println("PreHandler ping")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping .."))
	if err != nil {
		fmt.Println("PreHander fail")
	}
}

// Handle ..
func (pingRouter *PingRouter) Handle(request xinterface.IRequest) {
	fmt.Println("Handler ping")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping .."))
	if err != nil {
		fmt.Println("Hander fail")
	}
}

// PostHandle ..
func (pingRouter *PingRouter) PostHandle(request xinterface.IRequest) {
	fmt.Println("PostHandler ping")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping .."))
	if err != nil {
		fmt.Println("PostHander fail")
	}
}

func main() {
	s := xnet.NewServer("zinx V0.1")

	s.AddRouter(&PingRouter{})

	s.Serve()
}
