package root

import (
	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func Handle(ctx *types.Context) error {
	utils.WriteResponse(ctx.Conn, types.Response{
		StatusCode: config.OK,
	})
	return nil
}
