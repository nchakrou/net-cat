package funces

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

const maxip = 10

var msgs []string

var addrs = make(map[net.Conn]bool)
var names = make(map[net.Conn]string)
var mu sync.Mutex
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
	"     `-'       `--'`\n" +
	"[ENTER YOUR NAME]:"

func HandleConnection(conn net.Conn) {
	mu.Lock()
	if len(addrs) >= maxip {
		mu.Unlock()
		fmt.Println("Max IPs reached.")
		return
	}
	addrs[conn] = true
	mu.Unlock()

	conn.Write([]byte(Welcome))

	scanner := bufio.NewScanner(conn)
	message := ""
	for scanner.Scan() {
		is := true
		username := scanner.Text()
		username = strings.TrimSpace(username)
		mu.Lock()
		if username != "" {
			for _, v := range names {
				if username == v {
					is = false
					conn.Write([]byte("this name is aready taken\n"))
					conn.Write([]byte("[ENTER YOUR NAME]:"))
					break
				}
			}
		} else {
			is = false
			conn.Write([]byte("[ENTER YOUR NAME]:"))
		}
		if is {
			names[conn] = username
			message = username + " has joined our chat...\n"
			mu.Unlock()
			break
		}
		mu.Unlock()
	}
	mu.Lock()
	for v, _ := range addrs {
		if v != conn && names[v] != "" {
			v.Write([]byte(message))
		} else if conn == v {
			for _, k := range msgs {
				v.Write([]byte(k))
			}
		}
	}
	mu.Unlock()

	scannerr := bufio.NewScanner(conn)
	for scannerr.Scan() {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		msg := scannerr.Text()
		if msg == "" {
			continue
		}
		mu.Lock()
		for v, _ := range addrs {
			if v != conn && names[v] != "" {
				final := fmt.Sprintf("[%s][%s]:[%s]\n", currentTime, names[conn], msg)
				msgs = append(msgs, final)
				v.Write([]byte(final))
			}
		}
		mu.Unlock()
	}
	mu.Lock()
	if names[conn] != "" {
		disconnectMsg := names[conn] + " has left our chat...\n"
		for v := range addrs {
			if v != conn {
				v.Write([]byte(disconnectMsg))
			}
		}
		delete(addrs, conn)
		delete(names, conn)
	}
	mu.Unlock()
	conn.Close()
}
