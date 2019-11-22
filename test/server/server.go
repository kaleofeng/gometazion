package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":10080")
	if err != nil {
		fmt.Printf("Net - Listen: err(%v)\n", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Listener - Accept: err(%v)\n", err)
			break
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 64)
	for {
		n, err := conn.Read(buf)
		fmt.Printf("Conn - Read: buf(%s) n(%d) err(%v)\n", string(buf), n, err)
		if 0 == n || err != nil {
			break
		}
		conn.Write(buf)
	}
}
