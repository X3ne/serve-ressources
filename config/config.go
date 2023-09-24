package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTP	HTTPConfig
	CDN		CDNConfig
	REDIS	REDISConfig
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	return &Config{
		HTTP: LoadHTTPConfig(),
		CDN:  LoadCDNConfig(),
		REDIS: LoadREDISConfig(),
	}
}
