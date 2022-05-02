package broker

import (
	"fmt"
	"net"
)

func Run() {
	listen, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		fmt.Print("listen failed, err:", err)
		return
	}
	go Save()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Print("accept failed, err:", err)
			continue
		}
		go Process(conn)

	}
}
