package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"net"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
)

func ParseRequest(conn net.Conn) (*types.Request, error) {
	reader := bufio.NewReader(conn)

	// Read the request line first
	reqLine, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fmt.Println(reqLine)
	reqLine = strings.Trim(reqLine, config.CRLF)
	fmt.Println(reqLine)
	parts := strings.Split(reqLine, " ")
	if len(parts) < 2 {
		return nil, fmt.Errorf("malformed request line")
	}
	headers := make(map[string]string)
	//  Read headers , headers can be multliple so loop throught,
	var contentLength int64

	for {
		line, err := reader.ReadString('\n')
		if err != nil || line == config.CRLF {
			break
		}
		line = strings.TrimRight(line, "\r\n")
		if line == "" {
			break
			// end of header
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
			if strings.ToLower(key) == "content-length" {
				contentLength, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					return nil, fmt.Errorf("failed to convert content length to int: %w", err)
				}

			}

		}
	}
	// 	Return io.LimitReader to ensure only the body gets read.

	var bodyReader io.Reader
	if contentLength > 0 {
		bodyReader = io.LimitReader(reader, contentLength)

	}

	return &types.Request{
		Method:        parts[0],
		Path:          parts[1],
		Headers:       headers,
		Body:          bodyReader,
		ContentLength: contentLength,
	}, nil
	// Optional head;

	// if val, ok := headers["Content-Length"]; ok {
	// 	length, _ := strconv.Atoi(val)
	// 	body := make([]byte, length)
	// 	io.ReadFull(reader, body)
	// 	// Now you can use `body`
	// }
}
