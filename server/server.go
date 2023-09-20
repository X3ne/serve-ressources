package server

import (
	"serve-ressources/config"
	"serve-ressources/db"
	"serve-ressources/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	Echo		*echo.Echo
	DB			*gorm.DB
	REDIS		*services.RedisService
	Config	*config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Echo:		echo.New(),
		DB:			db.Init(cfg),
		REDIS:	services.Init(cfg),
		Config:	cfg,
	}
}

func (s *Server) Start(port string) error {
	return s.Echo.Start(":" + port)
}
