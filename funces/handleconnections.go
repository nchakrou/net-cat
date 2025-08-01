package funces

import (
	"bufio"
	"fmt"
	"net"
)

const maxip = 10

var addrs = make(map[net.Conn]bool)
var names = make(map[net.Conn]string)
var Welcome = "Welcome to TCP-Chat!\n" +
	"         _nnnn_\n" +
	"        dGGGGMMb\n" +
	"       @p~qp~~qMb\n" +
	"       M|@||@) M|\n" +
	"       @,----.JM|\n" +
	"      JS^\\__/  qKL\n" +
	"     dZP        qKRb\n" +
	"    dZP          qKKb\n" +
	"   fZP            SMMb\n" +
	"   HZM            MMMM\n" +
	"   FqM            MMMM\n" +
	" __| \".        |\\dS\"qML\n" +
	" |    `.       | `' \\Zq\n" +
	"_)      \\.___.,|     .'\n" +
	"\\____   )MMMMMP|   .'\n" +
	"     `-'       `--'`\n"

func HandleConnection(conn net.Conn) {

	if len(addrs) == maxip {
		fmt.Println("Max IPs reached.")
		return
	}
	addrs[conn] = true

	conn.Write([]byte(Welcome))

	scanner := bufio.NewScanner(conn)
	message := ""
	if scanner.Scan() {
		username := scanner.Text()
		names[conn] = username
		message = "the " + username + " has joined\n"

	}
	for v, _ := range addrs {
		if v != conn && Isvalid(names, v) {
			v.Write([]byte(message))
		}

	}

}
func Isvalid(addrs map[net.Conn]string, v net.Conn) bool {
	for k, _ := range addrs {
		if k == v {
			return true
		}

	}
	return false
}
