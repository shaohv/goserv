package xinterface

import (
	"net"
)

//IConnction define conn
type IConnection interface {
	// Start conn, let conn work now
	Start()

	// Stop conn
	Stop()

	// GetTCPConnection 从当前连接获取原始的socket TCPconn
	GetTCPConnection() *net.TCPConn

	// GetConnID()
	GetConnID() uint32

	// RemoteAddr
	RemoteAddr() net.Addr

	SendMsg(msgID uint32, data []byte) error
}

// HandFunc define a universal handle
type HandFunc func(*net.TCPConn, []byte, int) error
