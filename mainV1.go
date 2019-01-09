package main

import (
	"fmt"
	"net"
	"strings"
	"xudp"
)

func main() {
	addr := &net.UDPAddr{
		Port: 8888,
		IP:   net.ParseIP("10.1.0.54"),
	}
	server := xudp.NewServer(addr, func(req xudp.HandleRequest) {
		index := strings.Index(req.Data, "|")
		if index == -1 {
			fmt.Println(`69u6f24byh`, xudp.GetClientIdWithAddr(req.SenderInfo.Addr))
			return
		}
		toId := req.Data[:index]
		realData := req.Data[index:]
		err := req.Server.SendDataToClient(toId, realData)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	server.Run()
}
