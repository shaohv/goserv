package xnet

import (
	"fmt"
	"goserv/xinterface"
	"net"
)

// Connection represents a conn
type Connection struct {
	//Tcp socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	//handleAPI xinterface.HandFunc
	Router xinterface.IRouter

	ExitBuffChan chan bool
}

// NewConnection use to new a conn
// func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI xinterface.HandFunc) *Connection {
// 	c := &Connection{
// 		Conn:         conn,
// 		ConnID:       connID,
// 		isClosed:     false,
// 		handleAPI:    callbackAPI,
// 		ExitBuffChan: make(chan bool, 1),
// 	}

// 	return c
// }

// NewConnection use to new a conn
func NewConnection(conn *net.TCPConn, connID uint32, router xinterface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

// StartReader read data and dispatch request
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroute is start")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit! ")

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read from coonn fail ", err)
			c.ExitBuffChan <- true
			continue
		}

		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("handle fail")
		// 	c.ExitBuffChan <- true
		// 	return
		// }
		req := Request{
			conn: c,
			data: buf,
		}
		fmt.Println("11111")
		go func(request xinterface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// Start setup a goroute for read data and wait the goroute exit
func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			fmt.Println("goroute exit")
			return
		}
	}
}

// Stop close tcpconn and channel
func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	c.Conn.Close()

	c.ExitBuffChan <- true

	close(c.ExitBuffChan)
}

// GetTCPConnection get tcpconn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID as is name
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr as is name
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
