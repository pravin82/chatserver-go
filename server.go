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
	for cmd := range s.commands {
		switch cmd.id {
		case CMD_NAME:
			s.name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.listRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)

		case CMD_QUIT:
			s.quit(cmd.client, cmd.args)

		}
	}
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

func (s *server) name(c *client, args []string) {
	c.name = args[1]
	c.msg(fmt.Sprintf("Ok you have been assigned %s name", c.name))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r

	}
	r.members[c.conn.RemoteAddr()] = c
	s.quitCurrentRoom(c)
	c.room = r
	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("welcome to %s", r.name))

}

func (s *server) listRooms(c *client, args []string) {

}

func (s *server) msg(c *client, args []string) {

}

func (s *server) quit(c *client, args []string) {

}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s has left the room", c.name))
	}
}
