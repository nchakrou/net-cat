package funces

import (
	"fmt"
	"net"
)

func HostServer(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go HandleConnection(conn)
	}

}
