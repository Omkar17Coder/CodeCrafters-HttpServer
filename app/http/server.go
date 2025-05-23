package http

import (
	"flag"
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/http/handlers/echo"
	"github.com/codecrafters-io/http-server-starter-go/app/http/handlers/files"
	"github.com/codecrafters-io/http-server-starter-go/app/http/handlers/root"
	"github.com/codecrafters-io/http-server-starter-go/app/http/handlers/user_agent"
	"github.com/codecrafters-io/http-server-starter-go/app/types"
	"github.com/codecrafters-io/http-server-starter-go/app/utils"
)

var dirFlag *string

type Server struct {
	router *Router
}

func NewServer() *Server {
	server := &Server{
		router: NewRouter(),
	}

	// Register routes
	server.router.Register("GET", "/", root.Handle)
	server.router.Register("GET", "/echo/*", echo.Handle)
	server.router.Register("GET", "/user-agent", user_agent.Handle)
	server.router.Register("GET", "/files/*", files.GetFile)
	server.router.Register("POST", "/files/*", files.SetFile)
	server.router.PrintTree()

	return server
}

func resolveBaseDiretory() (string, error) {
	dirFlag = flag.String("directory", "", "Base directory where files stored")
	flag.Parse()

	if *dirFlag == "" {
		fmt.Println("provide the directory")
		return "", fmt.Errorf("please provide directory using --directory")
	}
	currDir, err := os.Getwd()
	fmt.Println(currDir)

	if err != nil {
		fmt.Printf("Failed to resolve path: %v", err.Error())
		return "", fmt.Errorf("failed to resolve path")
	}

	return currDir, nil
}

func StartServer() {
	baseDir, err := resolveBaseDiretory()
	if err != nil {
		fmt.Println("Failed with error:", err.Error())
		return
	}
	blue := "\033[34m"
	reset := "\033[0m"
	fmt.Println("Sever starting on port 4221..")
	fmt.Print(blue + `
 ██████╗ ███╗   ███╗██╗  ██╗ █████╗ ██████╗ 
██╔═══██╗████╗ ████║██║ ██╔╝██╔══██╗██╔══██╗
██║   ██║██╔████╔██║█████╔╝ ███████║██████╔╝
██║   ██║██║╚██╔╝██║██╔═██╗ ██╔══██║██╔══██╗
╚██████╔╝██║ ╚═╝ ██║██║  ██╗██║  ██║██║  ██║
 ╚═════╝ ╚═╝     ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝
                                            
` + reset)

	listner, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("failed to bind to port 4221 ", err.Error())
		os.Exit(1)
	}

	server := NewServer()
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go server.HandleConnection(conn, baseDir)
	}
}

func (s *Server) HandleConnection(conn net.Conn, baseDir string) {
	defer conn.Close()
	req, err := utils.ParseRequest(conn)
	if err != nil {
		utils.WriteResponse(conn, types.Response{
			StatusCode: 400,
		})
		return
	}

	ctx := &types.Context{
		Conn:    conn,
		Request: req,
		BaseDir: baseDir,
	}

	handler, found := s.router.Route(req.Method, req.Path)

	if !found {
		utils.WriteResponse(conn, types.Response{
			StatusCode: 404,
		})
		return
	}

	if err := handler(ctx); err != nil {
		utils.WriteResponse(conn, types.Response{
			StatusCode: 400,
		})
		return
	}
}
