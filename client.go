package chatserver_go

import "net"

type client struct {
	conn     net.Addr
	name     string
	room     *room
	commands chan<- command
}
