package main

import (
	"fmt"
	app "serve-ressources"
	"serve-ressources/config"

	"serve-ressources/docs"
)

func main() {
	cfg := config.NewConfig()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)

	app.Start(cfg)
}
