package xnet

import (
	"errors"
	"fmt"
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

			dealConn := NewConnection(conn, cid, CallBackToClient)

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

	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}