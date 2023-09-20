package app

import (
	"serve-ressources/config"
	"serve-ressources/server"
	"serve-ressources/server/v1/routes"

	"log"
)


func Start(cfg *config.Config) {
	app := server.NewServer(cfg)

	routes.ConfigureV1Routes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
