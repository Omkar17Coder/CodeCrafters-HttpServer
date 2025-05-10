package files

import (
	"fmt"
	"io"
	"os"
	
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func SetFile(ctx *types.Context) error {
	tokens := strings.Split(ctx.Request.Path, "/")
	if len(tokens) < 3 {
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return fmt.Errorf("failed to get enough data")
	}
	fileName := tokens[2]
	finalPath := fmt.Sprintf("%s/app/tmp/%s", ctx.BaseDir, fileName)

	if ctx.Request.ContentLength <= 0 || ctx.Request.Body == nil {
		return fmt.Errorf("invalid content length or missing body")
	}

	const maxInMemory = 1 << 20
	tempPath := finalPath + ".tmp"

	tmpFile, err := os.OpenFile(tempPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer tmpFile.Close()

	_, err = io.CopyN(tmpFile, ctx.Request.Body, ctx.Request.ContentLength)
	if err != nil && err != io.EOF {
		return fmt.Errorf("failed to stream to temp file: %w", err)
	}

	err = os.Rename(tempPath, finalPath)
	if err != nil {
		return fmt.Errorf("failed to rename temp file to final: %w", err)
	}
	fmt.Println(finalPath)
	utils.WriteResponse(ctx.Conn, types.Response{
		StatusCode: config.OK,
		Body:       "Created",
	})

	return nil
}
