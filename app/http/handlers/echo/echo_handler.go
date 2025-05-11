package echo

import (
	"strings"

	"fmt"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func Handle(ctx *types.Context) error {
	tokens := strings.Split(ctx.Request.Path, "/")
	if len(tokens) < 3 {
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return nil
	}

	text := tokens[2]
	res := types.Response{
		StatusCode: config.OK,
		Body:       text,
		Headers: map[string]string{
			"Content-Type":   config.TextContentType,
			"Content-Length": fmt.Sprintf("%d", len(text)),
		},
	}
	res = utils.CompressResponse(res, *ctx.Request)

	utils.WriteResponse(ctx.Conn, res)
	return nil
}
