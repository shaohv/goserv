package main

import (
	"fmt"
	"goserv/xnet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Client1 Test ... start")

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("Client Dial err ", err)
		return
	}

	dp := xnet.NewDataPack()
	for {

		sendData, _ := dp.Pack(xnet.NewMsgPkg(2, []byte("hello	ZINX from client 1")))

		_, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("client write err ", err)
			return
		}

		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
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

		//fmt.Printf("Server echo back: %s, cnt = %d\n", buf, cnt)

		time.Sleep(3 * time.Second)
	}
}
