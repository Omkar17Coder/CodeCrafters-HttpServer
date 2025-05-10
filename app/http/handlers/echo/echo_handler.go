package echo

import (
	"strings"

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
	utils.WriteResponse(ctx.Conn, types.Response{
		StatusCode: config.OK,
		Body:       text,
	})
	return nil
}
