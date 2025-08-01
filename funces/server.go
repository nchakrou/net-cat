package funces

import "net"

func HostServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		//handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
		}
		go HandleConnection(conn)
	}

}
