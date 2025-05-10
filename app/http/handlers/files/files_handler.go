package files

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/pkg/config"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

func GetFile(ctx *types.Context) error {

	tokens := strings.Split(ctx.Request.Path, "/")
	if len(tokens) < 3 {
	
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return fmt.Errorf("invalid path")
	}

	filePath := tokens[2]
	actualFile := fmt.Sprintf("%s/app/tmp/%s", ctx.BaseDir, filePath)
	
	file, err := os.Open(actualFile)
	if err != nil {
		if os.IsNotExist(err) {
			utils.WriteResponse(ctx.Conn, types.Response{
				StatusCode: config.NotFound,
			})
			return fmt.Errorf("file not found")
		}
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	data := make([]byte, 1024)
	count, err := file.Read(data)
	if err != nil {
		utils.WriteResponse(ctx.Conn, types.Response{
			StatusCode: config.BadRequest,
		})
		return fmt.Errorf("failed to read file: %w", err)
	}

	fileData := string(data[:count])
	utils.WriteResponse(ctx.Conn, types.Response{
		StatusCode: config.OK,
		Body:       fileData,
		Headers: map[string]string{
			"Content-Type":   config.FileContentType,
			"Content-Length": fmt.Sprintf("%d", count),
		},
	})

	return nil
}
