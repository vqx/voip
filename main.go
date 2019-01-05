package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	lister, err := net.Listen("tcp4", "10.1.0.41:8888")
	if err != nil {
		fmt.Println(1, err)
		return
	}
	for {
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println(2, err)
		} else {
			go func() {
				for {
					conn.Write([]byte("I'm alive"))
					time.Sleep(time.Second * 2)
				}
			}()
			go func() {
				for {
					data := make([]byte, 2048)
					n, err := conn.Read(data)
					if err != nil {
						fmt.Println(2, err)
						break
					}
					if n != 0 {
						fmt.Println(conn.RemoteAddr(), "get data success", string(data[:n]))
						conn.Write([]byte("server get data success"))
					}
					time.Sleep(time.Nanosecond * 100)
				}
			}()
		}
	}
}
