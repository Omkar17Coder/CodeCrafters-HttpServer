package echo

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func Handle(ctx *types.Context) {
	tokens := strings.Split(ctx.Request.Path, "/")
	if len(tokens) < 3 {
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return
	}

	echoBody := tokens[2]+"\n"
	utils.WriteResponse(ctx.Conn, types.Response{
		StatusCode: config.OK,
		Body:       echoBody,
		Headers: map[string]string{
			"Content-Type":   config.TextContentType,
			"Content-Length": fmt.Sprintf("%d", len(echoBody)),
		},
	})
}
