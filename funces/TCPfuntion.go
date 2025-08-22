package funces

import (
	"fmt"
	"net"
)

func Connection(port string) {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error of starting server", err)
		return
	}
	fmt.Println("the server start listen in the port " + port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error The accipting connection:", err)
			return
		}
		go HandleConnection(conn)
	}
}
