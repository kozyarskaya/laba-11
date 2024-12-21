package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Server  *echo.Echo
	address string

	uc Usecase
}

func NewServer(ip string, port int, uc Usecase) *Server {
	api := Server{
		uc: uc,
	}
	api.Server = echo.New()
	api.Server.POST("/reg", api.signUp)
	api.Server.POST("/aui", api.signIn)
	//api.Server.Use(JWTMiddleware)
	api.address = fmt.Sprintf("%s:%d", ip, port)

	return &api
}

func (api *Server) Run() {
	api.Server.Logger.Fatal(api.Server.Start(api.address))
}
