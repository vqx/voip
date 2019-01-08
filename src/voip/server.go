package voip

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	Listener  net.Listener
	ClientMap map[string]*ClientInfo
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", "10.1.0.54:8888")
	if err != nil {
		panic(`f4gsfejw3r ` + err.Error())
	}
	s.Listener = listen
	for {
		conn, err := listen.Accept()
		if err != nil {

		}
		if s.ClientMap == nil {
			s.ClientMap = map[string]*ClientInfo{}
		}
		client := &ClientInfo{
			Conn: conn,
		}
		id, err := client.getClientId()
		if err != nil {
			fmt.Println(`e6pqxuqg5a`, err.Error())
			continue
		}
		s.ClientMap[id] = client
	}
}

const (
	ClientStatusOnline = "Online"
	ClientStatusDail   = "Dail"
	ClientStatusBusy   = "Busy"
)

type ClientInfo struct {
	Id     string //ip and port hex, FF FF FF FF FF FF
	Status string
	Conn   net.Conn
	server *Server
}

func (c *ClientInfo) Dail() {

}
func (c *ClientInfo) getClientId() (string, error) {
	addr := c.Conn.RemoteAddr().String()
	if addr == "" {
		return "", errors.New(`jrdc4q4ttt addr == ""`)
	}
	tmp := strings.Split(addr, ":")
	var result string
	ipTmp := strings.Split(tmp[0], ".")
	for _, item := range ipTmp {
		partInt, err := strconv.Atoi(item)
		if err != nil {
			return "", errors.New(`vgbt9ftq5y err != nil ` + err.Error())
		}
		part := strconv.FormatInt(int64(partInt), 16)
		if len(part) == 1 {
			part = "0" + part
		}
		result += part
	}
	port, err := strconv.Atoi(tmp[1])
	if err != nil {
		return "", errors.New(`ex2caz94pv err != nil ` + err.Error())
	}
	portPart := strconv.FormatInt(int64(port), 16)
	for len(portPart) < 4 {
		portPart = "0" + portPart
	}
	result += portPart
	return result, nil
}
