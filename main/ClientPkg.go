package main

import (
	"fmt"
	"goserv/xnet"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("dail fail ", err)
		return
	}

	dp := xnet.NewDataPack()

	msg1 := &xnet.Message{
		Id:      1,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendData1, err := dp.Pack(msg1)
	if err != nil {
		return
	}

	msg2 := &xnet.Message{
		Id:      2,
		DataLen: 5,
		Data:    []byte{'w', 'o', 'r', 'l', 'd'},
	}

	sendData2, err := dp.Pack(msg2)
	if err != nil {
		return
	}

	sendData1 = append(sendData1, sendData2...)

	conn.Write(sendData1)

	select {}

}
