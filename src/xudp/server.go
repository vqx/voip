package xudp

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

func NewServer(addr *net.UDPAddr, handle func(HandleRequest)) (s *Server) {
	return &Server{
		serverAddr: addr,
		ClientMap:  map[string]*ClientInfo{},
		Handle:     handle,
	}
}

type Server struct {
	serverAddr *net.UDPAddr
	ServerConn *net.UDPConn
	ClientMap  map[string]*ClientInfo
	clientLock sync.Mutex
	Handle     func(HandleRequest)
	closeFlag  bool
}

func (s *Server) Run() {
	conn, err := net.ListenUDP("udp4", s.serverAddr)
	if err != nil {
		panic(`f4gsfejw3r ` + err.Error())
	}
	fmt.Println(`te6jkmkvus`, `udp server run in`, s.serverAddr.IP.String()+":"+strconv.Itoa(s.serverAddr.Port))
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
	for !s.closeFlag {
		time.Sleep(time.Second)
	}
}

func (s *Server) HandleData(addr *net.UDPAddr, data []byte) {
	id := GetClientId(addr.IP.String(), addr.Port)
	sender, ok := s.ClientMap[id]
	if !ok {
		fmt.Println(`xv3wqwdgw3`, `new client`, len(s.ClientMap)+1)
		sender = &ClientInfo{
			Id:     id,
			Addr:   addr,
			server: s,
		}
		s.ClientMap[id] = sender
	}
	if s.Handle != nil {
		s.Handle(HandleRequest{
			Server:     s,
			SenderInfo: sender,
			Data:       string(data),
		})
	}
}

func (s *Server) SendDataToClient(id string, data string) error {
	toClientInfo, ok := s.ClientMap[id]
	if !ok {
		errMsg := `6gxn6uer7p client ` + id + " no online"
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}
	s.clientLock.Lock()
	n, err := s.ServerConn.WriteToUDP([]byte(data), toClientInfo.Addr)
	s.clientLock.Unlock()
	if err != nil {
		errMsg := `m9y56dfvde WriteToUDP ` + id + " " + err.Error()
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}
	if n == 0 {
		errMsg := `5yqqwe7zmf WriteToUDP ` + id + " nothing to send"
		fmt.Println(errMsg)
		return errors.New(errMsg)
	}
	//fmt.Println(`yb9rgr57as`, "WriteToUDP", id, "to", getId(req.SenderInfo.Addr), realData)
	return nil
}

func (s *Server) Close() {
	//s.clientLock.Lock()
	s.closeFlag = true
	//s.clientLock.Unlock()
}

func GetClientId(ip string, port int) string {
	var result string
	ipTmp := strings.Split(ip, ".")
	for _, item := range ipTmp {
		partInt, err := strconv.Atoi(item)
		if err != nil {
			partInt = 0
		}
		part := strconv.FormatInt(int64(partInt), 16)
		if len(part) == 1 {
			part = "0" + part
		}
		result += part
	}
	portPart := strconv.FormatInt(int64(port), 16)
	for len(portPart) < 4 {
		portPart = "0" + portPart
	}
	result += portPart
	return result
}

func GetClientIdWithAddr(addr *net.UDPAddr) string {
	return GetClientId(addr.IP.String(), addr.Port)
}

type HandleRequest struct {
	Server     *Server
	SenderInfo *ClientInfo
	Data       string
}
