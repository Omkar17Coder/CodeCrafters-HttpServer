package types

import "io"

type Request struct {
	Method  string
	Path    string
	Headers map[string]string
	Body    io.Reader
	ContentLength int64
}
