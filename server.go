package main

import (
	"fmt"
	"net"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {

}

func (s *server) newClient(conn net.Conn) {
	fmt.Println("New client had connneted %s", conn.RemoteAddr())
	client := &client{
		conn:     conn,
		name:     "anon",
		commands: s.commands,
	}
	client.readInput()
}
