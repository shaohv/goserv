package xnet

import (
	"errors"
	"fmt"
	"goserv/util"
	"goserv/xinterface"
	"io"
	"net"
)

// Connection represents a conn
type Connection struct {
	TcpServer xinterface.IServer
	//Tcp socket
	Conn *net.TCPConn

	ConnID uint32

	isClosed bool

	//handleAPI xinterface.HandFunc
	//Router xinterface.IRouter
	MsgHandle xinterface.IMsgHandle

	ExitBuffChan chan bool

	msgChan chan []byte
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
func NewConnection(server xinterface.IServer, conn *net.TCPConn, connID uint32,
	msgHandle xinterface.IMsgHandle) *Connection {
	c := &Connection{
		TcpServer: server,
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		//Router:       router,
		MsgHandle:    msgHandle,
		ExitBuffChan: make(chan bool, 1),
		msgChan:      make(chan []byte),
	}

	c.TcpServer.GetConnMgr().Add(c)
	return c
}

// StartReader read data and dispatch request
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroute is start")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit! ")
	defer c.Stop()
	dp := NewDataPack()

	for {
		// buf := make([]byte, 512)
		// _, err := c.Conn.Read(buf)
		// if err != nil {
		// 	fmt.Println("read from coonn fail ", err)
		// 	c.ExitBuffChan <- true
		// 	continue
		// }

		// if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
		// 	fmt.Println("handle fail")
		// 	c.ExitBuffChan <- true
		// 	return
		// }

		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.Conn, headData)
		if err != nil {
			fmt.Println("read head fail ", err)
			break
		}

		msg, err := dp.UnPack(headData)
		if err != nil {
			fmt.Println("unpack head fail ", err)
			return
		}

		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(c.Conn, data)
			if err != nil {
				fmt.Println("read date fail ", err)
				return
			}

			fmt.Println("===> Rcv Msg ID=", msg.GetMsgId(), " DataLen=", msg.GetDataLen(), " data=", string(data))
		}
		msg.SetData(data)

		req := Request{
			conn: c,
			msg:  msg,
		}

		// fmt.Println("11111")
		// go func(request xinterface.IRequest) {
		// 	c.Router.PreHandle(request)
		// 	c.Router.Handle(request)
		// 	c.Router.PostHandle(request)
		// }(&req)

		if util.GlobalObject.WorkPoolSize > 0 {
			c.MsgHandle.SendMsgToTaskQueue(&req)
		} else {
			go c.MsgHandle.DoMsgHandle(&req)
		}
	}
}

// Start setup a goroute for read data and wait the goroute exit
func (c *Connection) Start() {
	go c.StartReader()

	go c.StartWriter()

	for {
		select {
		case <-c.ExitBuffChan:
			fmt.Println("goroute exit")
			return
		}
	}
}

// StartWriter ..
func (c *Connection) StartWriter() {
	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				c.isClosed = true
				return
			}
		case <-c.ExitBuffChan:
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

	c.TcpServer.GetConnMgr().Remove(c)

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

// SendMsg ..
func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	dp := NewDataPack()

	if c.isClosed == true {
		return errors.New("conn is closed")
	}

	sendData, err := dp.Pack(NewMsgPkg(msgID, data))
	if err != nil {
		return errors.New("pack msg fail")
	}

	c.msgChan <- sendData

	// if _, err := c.Conn.Write(sendData); err != nil {
	// 	c.isClosed = true
	// 	return err
	// }

	return nil
}
