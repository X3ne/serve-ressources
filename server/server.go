package server

import (
	"serve-ressources/config"
	"serve-ressources/services"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo		*echo.Echo
	REDIS		*services.RedisService
	Config	*config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:		echo.New(),
		REDIS:	services.Init(cfg),
		Config:	cfg,
	}
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}
