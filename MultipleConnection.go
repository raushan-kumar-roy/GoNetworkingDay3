package main

import (
	"fmt"
	"net"
	"sync"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Handling connection from", conn.RemoteAddr())
	conn.Close()
	fmt.Println("Closed connection from", conn.RemoteAddr())
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on 0.0.0.0:8080")

	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go handleConnection(conn, &wg)
	}
	wg.Wait()
}
