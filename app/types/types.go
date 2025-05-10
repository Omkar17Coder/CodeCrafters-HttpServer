package types

import (
	"net"
)

type HandlerFunc func(*Context) error

type Context struct {
	Conn    net.Conn
	Request *Request
	BaseDir string
}
