package funces

import (
	"fmt"
	"net"
	"time"
)

func MessageForma(conn net.Conn) string {
	time := time.Now().Format("2006-01-02 15:04:05")
	format := fmt.Sprintf("[%s][%s]:", time, names[conn])
	return format
}
