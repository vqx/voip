package voip

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Server struct {
	serverAddr *net.UDPAddr
	ServerConn *net.UDPConn
	ClientMap  map[string]*ClientInfo
	clientLock sync.Mutex
}

func (s *Server) Run() {
	addr := net.UDPAddr{
		Port: 8888,
		IP:   net.ParseIP("10.1.0.54"),
	}
	conn, err := net.ListenUDP("udp4", &addr)
	if err != nil {
		panic(`f4gsfejw3r ` + err.Error())
	}
	s.ServerConn = conn
	go func() {
		for {
			data := make([]byte, 2048)
			n, remoteAddr, err := conn.ReadFromUDP(data)
			if err != nil {
				time.Sleep(1)
				fmt.Println("p4zu7zgexf", "ReadFromUDP", err)
				continue
			}
			if n == 0 {
				time.Sleep(1)
				continue
			}
			go s.HandleData(remoteAddr, data)
		}
	}()

	//for {
	//	n, remoteaddr, err := conn.ReadFromUDP(p)
	//	if err != nil {
	//
	//	}
	//	if s.ClientMap == nil {
	//		s.ClientMap = map[string]*ClientInfo{}
	//	}
	//	client := &ClientInfo{
	//		Addr: remoteaddr,
	//	}
	//	id, err := client.getClientId()
	//	if err != nil {
	//		fmt.Println(`e6pqxuqg5a`, err.Error())
	//		continue
	//	}
	//	s.ClientMap[id] = client
	//}
}

func (s *Server) HandleData(addr *net.UDPAddr, data []byte) {

}

const (
	ClientStatusOnline = "Online"
	ClientStatusDail   = "Dail"
	ClientStatusBusy   = "Busy"
)

type ClientInfo struct {
	Id     string //ip and port hex, FF FF FF FF FF FF
	Status string
	Addr   *net.UDPAddr
	server *Server
}

//func (c *ClientInfo) Dail() {
//
//}

func getClientId(addr *net.UDPAddr) (string, error) {
	ip := addr.IP.String()
	if ip == "" {
		return "", errors.New(`jrdc4q4ttt addr == ""`)
	}
	var result string
	ipTmp := strings.Split(ip, ".")
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
	portPart := strconv.FormatInt(int64(addr.Port), 16)
	for len(portPart) < 4 {
		portPart = "0" + portPart
	}
	result += portPart
	return result, nil
}

func (c *ClientInfo) getClientId() (string, error) {
	addr := c.Addr.IP.String()
	if addr == "" {
		return "", errors.New(`jrdc4q4ttt addr == ""`)
	}
	var result string
	ipTmp := strings.Split(addr, ".")
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
	portPart := strconv.FormatInt(int64(c.Addr.Port), 16)
	for len(portPart) < 4 {
		portPart = "0" + portPart
	}
	result += portPart
	return result, nil
}
