package utils

import (
	"fmt"
	"strings"

	csd "github.com/codecrafters-io/http-server-starter-go/app/pkg"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
)

func CompressResponse(res types.Response, req types.Request) types.Response {
	if strings.Contains(req.Headers["Accept-Encoding"], "gzip") {
		compressed, err := csd.GzipBytes([]byte(res.Body))
		if err != nil {
			return types.Response{
				StatusCode: 500,
			}
		}
		res.Body = string(compressed[:])

		res.Headers["Content-Encoding"] = "gzip"
	}
	res.Headers["Content-Length"] = fmt.Sprintf("%d", len(res.Body))
	return res

}
