package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	conn, err := net.Dial("tcp", "localhost:10080")
	if err != nil {
		fmt.Printf("Net - Dial: err(%v)\n", err)
		return
	}

	defer conn.Close()

	wg.Add(1)
	go handleConnection(conn)
	wg.Wait()
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 64)
	copy(buf, []byte("hello, server"))
	for {
		n, err := conn.Write(buf)
		fmt.Printf("Conn - Write: buf(%s) n(%d) err(%v)\n", string(buf), n, err)
		if 0 == n || err != nil {
			break
		}

		n, err = conn.Read(buf)
		fmt.Printf("Conn - Read: buf(%v) n(%d) err(%v)\n", buf, n, err)
		if 0 == n || err != nil {
			break
		}

		time.Sleep(time.Second * 5)
	}

	wg.Done()
}
