package services

import (
	"context"
	"fmt"
	"serve-ressources/config"

	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client	*redis.Client
	Ctx			context.Context
}

func Init(cfg *config.Config) *RedisService {

	ctx := context.Background()

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.REDIS.Host, cfg.REDIS.Port),
		Password: cfg.REDIS.Password,
		DB:       cfg.REDIS.DB,
	})


	_, err := client.Ping(ctx).Result()

	if err != nil {
		panic(err)
	}

	return &RedisService{
		Client: client,
		Ctx: ctx,
	}
}
