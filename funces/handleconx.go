package funces

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	messages []string
	address  = make(map[net.Conn]bool)
	names    = make(map[net.Conn]string)
	Welcome  = "Welcome to TCP-Chat!\n" +
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
	err error
	mux sync.Mutex
)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	mux.Lock()
	if len(address) >= 10 {
		conn.Write([]byte("The group is full"))
		mux.Unlock()
		return
	}
	address[conn] = true
	mux.Unlock()

	defer func() {
		mux.Lock()
		delete(address, conn)
		mux.Unlock()
	}()

	_, err = conn.Write([]byte(Welcome))
	if err != nil {
		fmt.Println("Error of Writing connection", err)
		return
	}
	username := ""
	scan := bufio.NewScanner(conn)
	for scan.Scan() {
		username = scan.Text()
		username = strings.TrimSpace(username)

		if Isvalid(username) {
			if !Isvalidname(username) {
				_, err = conn.Write([]byte("Sorry this name is taken\n"))
				if err != nil {
					fmt.Println("Error of Writing connection", err)
					return
				}
				_, err = conn.Write([]byte("[ENTER YOUR NAME]:"))
				if err != nil {
					fmt.Println("Error of Writing connection", err)
					return
				}
			} else {

				mux.Lock()
				names[conn] = username
				mux.Unlock()
				break
			}
		} else {
			_, err = conn.Write([]byte("Invalid name\n"))
			if err != nil {
				fmt.Println("Error of Writing connection", err)
				return
			}
			_, err = conn.Write([]byte("[ENTER YOUR NAME]:"))
			if err != nil {
				fmt.Println("Error of Writing connection", err)
				return
			}
		}
	}
	defer func() {
		mux.Lock()
		delete(names, conn)
		mux.Unlock()
	}()
	mux.Lock()
	if names[conn] == "" {
		mux.Unlock()
		return
	}
	mux.Unlock()

	mux.Lock()
	for _, v := range messages {
		conn.Write([]byte(v))
	}
	mux.Unlock()

	mux.Lock()
	for connaddres := range address {
		if connaddres != conn && names[connaddres] != "" {
			_, err = connaddres.Write([]byte("\n" + username + " has joined our chat...\n"))
			if err != nil {
				fmt.Println("Error of Writing connection", err)
				mux.Unlock()
				return
			}
			_, err = connaddres.Write([]byte(MessageForma(connaddres)))
			if err != nil {
				fmt.Println("Error of Writing connection", err)
				mux.Unlock()
				return
			}
		}
	}

	mux.Unlock()

	format := MessageForma(conn)
	_, err = conn.Write([]byte(format))
	if err != nil {
		fmt.Println("Error of Writing connection", err)
		return
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		message = strings.TrimSpace(message)

		format1 := MessageForma(conn)
		if message != "" {
			messages = append(messages, format+message+"\n")
		}
		format = format1

		if message == "" {
			conn.Write([]byte(format1))
			continue
		}

		conn.Write([]byte(format1))

		format2 := MessageForma(conn)
		for connaddres := range address {
			if connaddres != conn && names[connaddres] != "" && Isvalidmassage(message) {
				_, err = connaddres.Write([]byte("\n" + format2 + message + "\n"))
				if err != nil {
					fmt.Println("Error of Writing connection", err)
					return
				}
				_, err = connaddres.Write([]byte(MessageForma(connaddres)))
				if err != nil {
					fmt.Println("Error of Writing connection", err)
					return
				}
			}
		}

	}

	mux.Lock()
	for connaddres := range address {
		if connaddres != conn && names[connaddres] != "" {

			_, err = connaddres.Write([]byte("\n" + names[conn] + " has left our chat...\n"))
			if err != nil {
				fmt.Println("Error of Writing connection", err)
				mux.Unlock()
				return
			}
			format := MessageForma(connaddres)
			_, err = connaddres.Write([]byte(format))
			if err != nil {
				fmt.Println("Erro", err)
				mux.Unlock()
				return
			}
		}
	}
	mux.Unlock()
}
