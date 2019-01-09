package xudp

import "net"

type ClientInfo struct {
	Id     string //ip and port hex, FF FF FF FF FF FF
	//Status string
	Addr   *net.UDPAddr
	server *Server
}
