package xinterface

import "net"

//IConnction define conn
type IConnction interface {
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
}

// HandFunc define a universal handle
type HandFunc func(*net.TCPConn, []byte, int) error
