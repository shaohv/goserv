package xnet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func ClientTest() {
	time.Sleep(3 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello	ZINX"))
		if err != nil {
			fmt.Println("client write err ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read err ", err)
			return
		}

		fmt.Printf("Server echo back: %s, cnt = %d\n", buf, cnt)

		time.Sleep(3 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("[zinx	V0.1]")
	go ClientTest()
	s.Serve()
}
