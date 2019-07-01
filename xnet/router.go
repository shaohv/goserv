package xnet

import (
	"fmt"
	"goserv/xinterface"
)

// BaseRouter ...
type BaseRouter struct {
}

// PreHandle ...
func (br *BaseRouter) PreHandle(req xinterface.IRequest) {
	fmt.Println("222")
}

// Handle ...
func (br *BaseRouter) Handle(req xinterface.IRequest) {}

// PostHandle ...
func (br *BaseRouter) PostHandle(req xinterface.IRequest) {}
