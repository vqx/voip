package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var all []net.Conn
	lister, err := net.Listen("tcp4", "10.1.0.54:8888")
	if err != nil {
		fmt.Println(1, err)
		return
	}
	for {
		fmt.Println("wait client...")
		conn, err := lister.Accept()
		if err != nil {
			fmt.Println(2, err)
		} else {
			all = append(all, conn)
			fmt.Println("accept client", len(all), conn.RemoteAddr())
			if len(all) == 2 {
				fmt.Println("end accept")
				break
			}
		}
	}
	conn0 := all[0]
	conn1 := all[1]
	var nn int
	var mm int
	var lock1 sync.Mutex
	var lock2 sync.Mutex
	go func() {
		fmt.Println("start 1 to 2")
		for {
			data := make([]byte, 2048)
			lock1.Lock()
			n, err := conn0.Read(data)
			if err != nil {
				lock1.Unlock()
				fmt.Println("conn0.Read", err)
				break
			}
			lock1.Unlock()
			if n != 0 {
				if nn < 100 {
					fmt.Println("conn0.Read success ")
				}
				nn++
				lock2.Lock()
				_, err = conn1.Write(data[:n])
				if err != nil {
					fmt.Println("conn1.Write", err)
					lock2.Unlock()
					break
				}
				lock2.Unlock()
			}
		}
	}()
	go func() {
		fmt.Println("start 2 to 1")
		for {
			data := make([]byte, 2048)
			lock2.Lock()
			n, err := conn1.Read(data)
			if err != nil {
				fmt.Println("conn1.Read", err)
				lock2.Unlock()
				break
			}
			lock2.Unlock()
			if n != 0 {
				if mm < 100 {
					fmt.Println("conn0.Read success ")
				}
				mm++
				lock1.Lock()
				_, err = conn0.Write(data[:n])
				if err != nil {
					fmt.Println("conn0.Write", err)
					lock1.Unlock()
					break
				}
				lock1.Unlock()
			}
		}
	}()
	for {
		time.Sleep(time.Second)
	}
}
