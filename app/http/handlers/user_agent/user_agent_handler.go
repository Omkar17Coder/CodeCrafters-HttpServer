package user_agent

import (
	"fmt"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func Handle(ctx *types.Context) error {
	userAgent := ctx.Request.Headers["User-Agent"]
	res := types.Response{
		StatusCode: config.OK,
		Body:       userAgent,
		Headers: map[string]string{
			"Content-Type":   config.TextContentType,
			"Content-Length": fmt.Sprintf("%d", len(userAgent)),
		},
	}
	res = utils.CompressResponse(res, *ctx.Request)

	utils.WriteResponse(ctx.Conn, res)
	return nil
}
