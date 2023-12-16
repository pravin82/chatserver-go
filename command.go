package chatserver_go

type commandId int

const (
	CMD_NAME commandId = iota
	CMD_JOIN
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)
