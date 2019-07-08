package xnet

import (
	"errors"
	"fmt"
	"goserv/main/util"
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

	Router xinterface.IRouter
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

		for {

			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept err ", err)
				continue
			}

			var cid uint32
			cid = 0

			dealConn := NewConnection(conn, cid, s.Router)

			go dealConn.Start()
		}
	}()
}

// Stop is impletion of IServer Stop method
func (s *Server) Stop() {

	fmt.Println("[STOP] Server stoping, Name", s.Name)

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
		Router:    nil,
	}

	return s
}

// AddRouter ..
func (s *Server) AddRouter(router xinterface.IRouter) {
	s.Router = router
	return
}
