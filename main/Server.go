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
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping .."))
	// if err != nil {
	// 	fmt.Println("PreHander fail")
	// }
}

// Handle ..
func (pingRouter *PingRouter) Handle(request xinterface.IRequest) {
	fmt.Println("Handler ping")
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping .."))
	err := request.GetConnection().SendMsg(1, []byte("ping ping .."))
	if err != nil {
		fmt.Println("Hander fail")
	}
}

// PostHandle ..
func (pingRouter *PingRouter) PostHandle(request xinterface.IRequest) {
	fmt.Println("PostHandler ping")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping .."))
	// if err != nil {
	// 	fmt.Println("PostHander fail")
	// }
}

// HelloRouter is a applic router
type HelloRouter struct {
	xnet.BaseRouter
}

// PreHandle ..
func (pingRouter *HelloRouter) PreHandle(request xinterface.IRequest) {
	fmt.Println("PreHandler hello")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping .."))
	// if err != nil {
	// 	fmt.Println("PreHander fail")
	// }
}

// Handle ..
func (pingRouter *HelloRouter) Handle(request xinterface.IRequest) {
	fmt.Println("Handler hello")
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping .."))
	err := request.GetConnection().SendMsg(2, []byte("hello hello .."))
	if err != nil {
		fmt.Println("Hander fail")
	}
}

// PostHandle ..
func (pingRouter *HelloRouter) PostHandle(request xinterface.IRequest) {
	fmt.Println("PostHandler hello")
	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping .."))
	// if err != nil {
	// 	fmt.Println("PostHander fail")
	// }
}

// DoConnectionBegin ...
func DoConnectionBegin(conn xinterface.IConnection) {
	fmt.Println("Conn begin...")
	err := conn.SendMsg(2, []byte("DO CONNECTION BEGIN ..."))
	if err != nil {
		fmt.Println("send conn begin msg fail!", err)
	}
}

// DoConnectionStop ...
func DoConnectionStop(conn xinterface.IConnection) {
	fmt.Println("Conn stop ...")
}

func main() {
	s := xnet.NewServer("zinx V0.1")

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionStop)
	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &HelloRouter{})

	s.Serve()
}
