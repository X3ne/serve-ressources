package config

import (
	"os"
	"strconv"
)

type REDISConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func LoadREDISConfig() REDISConfig {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	return REDISConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	}
}
