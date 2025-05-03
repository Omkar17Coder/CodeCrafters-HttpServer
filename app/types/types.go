package types

import (
	"net"
)



type HandlerFunc func(*Context)

type Context struct {
	Conn    net.Conn
	Request *Request
	BaseDir string
}
