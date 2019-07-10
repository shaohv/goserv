package util

import (
	"encoding/json"
	"fmt"
	"goserv/xinterface"
	"io/ioutil"
)

// GlobalObj store some global setting data
type GlobalObj struct {
	TCPServer xinterface.IServer

	Host string

	TCPPort int

	Name string

	Version string

	MaxPacketSize uint32

	MaxConn uint32
}

// GlobalObject ..
var GlobalObject *GlobalObj

// Reload ..
func (g *GlobalObj) Reload() {

	data, err := ioutil.ReadFile("conf/goserv.json")
	if err != nil {
		fmt.Println("read setting file fail ", err)
		panic(err)
	}

	err = json.Unmarshal(data, GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:          "goserv",
		Version:       "0.4",
		MaxConn:       1200,
		MaxPacketSize: 4096,
		TCPPort:       7777,
		Host:          "0.0.0.0",
	}

	GlobalObject.Reload()
}
