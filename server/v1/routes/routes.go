package routes

import (
	"fmt"
	s "serve-ressources/server"
	"serve-ressources/server/v1/handlers"

	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Torrents API
// @version 1.0.0
// @description v1.0.0 Torrents API

// @BasePath /v1
func ConfigureV1Routes(server *s.Server) {
	cdnHandler := handlers.NewRessourceHandler(server)

	v1 := server.Echo.Group("/v1")

	v1.Use(middleware.Logger())
	v1.Use(middleware.Recover())

	v1.GET("/docs/*", echoSwagger.WrapHandler)

	cdn := v1.Group("/cdn")
	cdn.GET("/:id", cdnHandler.GetRessource)

	fmt.Println("V1 routes configured")
}
