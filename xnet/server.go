package xnet

import (
	"errors"
	"fmt"
	"goserv/util"
	"goserv/xinterface"
	"net"
	"time"
)

// Server is implemention of IServer
type Server struct {
	Name string

	IPVersion string

	IP string

	Port int

	//Router xinterface.IRouter
	msgHandler xinterface.IMsgHandle

	ConnManager xinterface.IConnMgr

	OnCallStart func(conn xinterface.IConnection)

	OnCallStop func(conn xinterface.IConnection)
}

// CallBackToClient as request handle
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("Echo to client")

	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("Conn write fail ", err)
		return errors.New("Conn write fail")
	}

	return nil
}

// Start is implemention of IServer Start method
func (s *Server) Start() {

	if s == nil {
		fmt.Print("Receive nil param")
		return
	}

	fmt.Printf("[Start] Server listening, IP: %s, Port: %d", s.IP, s.Port)
	fmt.Printf("[Start] Server listening, IP: %s, Port: %d", util.GlobalObject.Host, util.GlobalObject.TCPPort)

	go func() {
		s.msgHandler.StartWorkerPool()

		// 1. construct tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2. listen
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("Start goserv ", s.Name, "success, now listenning")
		var cid uint32
		for {

			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ", err)
				continue
			}

			if uint32(s.GetConnMgr().Len()) >= util.GlobalObject.MaxConn {
				fmt.Println("Reach max conn, ", util.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			cid = cid + 1

			dealConn := NewConnection(s, conn, cid, s.msgHandler)

			go dealConn.Start()
		}
	}()
}

// Stop is impletion of IServer Stop method
func (s *Server) Stop() {

	fmt.Println("[STOP] Server stoping, Name", s.Name)

	s.GetConnMgr().ClearConn()
	// TODO Server.Stop()
}

//Serve now
func (s *Server) Serve() {

	s.Start()

	for {
		time.Sleep(10 * time.Second)
	}
}

// NewServer create a IServer handle
func NewServer(name string) xinterface.IServer {

	util.GlobalObject.Reload()

	s := &Server{
		Name:      util.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        util.GlobalObject.Host,
		Port:      util.GlobalObject.TCPPort,
		//Router:    nil,
		msgHandler:  NewMsgHandle(),
		ConnManager: NewConnMgr(),
	}

	return s
}

// AddRouter ..
func (s *Server) AddRouter(msgID uint32, router xinterface.IRouter) {
	//s.Router = router
	s.msgHandler.AddRouter(msgID, router)
	return
}

// GetConnMgr ..
func (s *Server) GetConnMgr() xinterface.IConnMgr {
	return s.ConnManager
}

// SetOnConnStart ..
func (s *Server) SetOnConnStart(hook func(xinterface.IConnection)) {
	s.OnCallStart = hook
}

// SetOnConnStop ..
func (s *Server) SetOnConnStop(hook func(xinterface.IConnection)) {
	s.OnCallStop = hook
}

// CallOnConnStart ..
func (s *Server) CallOnConnStart(conn xinterface.IConnection) {
	if s.OnCallStart != nil {
		fmt.Println("======")
		s.OnCallStart(conn)
	}
}

// CallOnConnStop ..
func (s *Server) CallOnConnStop(conn xinterface.IConnection) {
	if s.OnCallStart != nil {
		fmt.Println("-------")
		s.OnCallStop(conn)
	}
}
