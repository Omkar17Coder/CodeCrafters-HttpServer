package http

import (
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/modals"
)

type Context struct {
	Conn    net.Conn
	Req     *modals.Request
	BaseDir string
}
