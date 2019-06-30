package main

import (
	"goserv/xnet"
)

func main() {
	s := xnet.NewServer("zinx V0.1")

	s.Serve()
}
