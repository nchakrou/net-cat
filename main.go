package main

import (
	"fmt"
	"netcat/funces"
	"os"
)

func main() {
	args := os.Args[1:]
	port := "8989"
	if len(args) == 1 {
		port = args[0]
	} else if len(args) > 1 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}
	fmt.Println("Listening on the port :" + port)
	funces.HostServer(":" + port)
}
