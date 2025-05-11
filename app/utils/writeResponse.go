package utils

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
)

func WriteResponse(conn net.Conn, res types.Response) {
	statusText := map[int]string{
		config.OK:         "OK",
		config.NotFound:   "Not Found",
		config.BadRequest: "Bad Request",
	}

	if res.Headers == nil {
		res.Headers = make(map[string]string)
	}
	
	res.Headers["Content-Length"] = fmt.Sprintf("%d", len(res.Body))
	res.Headers["Content-Type"] = config.TextContentType
	res.Headers["Connection"] = "close"

	statusLine := fmt.Sprintf("%s %d %s\r\n", config.Server, res.StatusCode, statusText[res.StatusCode])

	headers := ""
	for k, v := range res.Headers {
		headers += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	final := statusLine + headers + config.CRLF + res.Body
	fmt.Println(final)
	conn.Write([]byte(final))
}
