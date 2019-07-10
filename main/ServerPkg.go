package main

import (
	"fmt"
	"goserv/xnet"
	"io"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept fail ", err)
			continue
		}

		go func(conn net.Conn) {
			dp := xnet.NewDataPack()

			for {
				headData := make([]byte, dp.GetHeadLen())
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head fail ", err)
					break
				}

				msgHead, err := dp.UnPack(headData)
				if err != nil {
					fmt.Println("unpack head fail ", err)
					return
				}

				if msgHead.GetDataLen() > 0 {
					msg := msgHead.(*xnet.Message)
					msg.Data = make([]byte, msg.GetDataLen())

					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("read date fail ", err)
						return
					}

					fmt.Println("===> Rcv Msg ID=", msg.GetMsgId(), " DataLen=", msg.GetDataLen(), " data=", string(msg.GetData()))
				}
			}

		}(conn)
	}
}
