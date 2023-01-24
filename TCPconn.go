package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8085")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	message := "Hello, World!"
	fmt.Fprintf(conn, message)

	conn.SetReadDeadline(time.Now().Add(time.Second))
	response := make([]byte, 4096)
	n, err := conn.Read(response)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(response[:n]))
}
