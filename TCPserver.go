package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8085")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening on :8080...")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	fmt.Printf("Received: %s", string(buf[:n]))
	conn.Write(buf[:n])
}
