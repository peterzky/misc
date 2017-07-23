package main

import (
	"fmt"
	"net"
	"os"
)

func SocketClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:3333")
	if err != nil {
		os.Exit(0)
	}

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	fmt.Println(string(buff[:n]))
}

func main() {
	SocketClient()
}
