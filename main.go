package main

import (
	"fmt"
	"net"
)

func main() {
	server := newServer()
	go server.run()
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Error in starting server")
		return
	}
	defer listener.Close()
	fmt.Println("Started server on :8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error in accepting connection")
			continue
		}
		go server.newClient(conn)

	}
}
