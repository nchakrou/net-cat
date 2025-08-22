package main

import (
	"fmt"
	"os"

	"net-cat/funces"
)

func main() {
	portDefult := ":8989"
	if len(os.Args) == 2 {
		portDefult = ":" + os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}
	funces.Connection(portDefult)
}
