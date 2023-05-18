package redis

import (
	"auth/configs"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	Redis *redis.Client
}

func NewRedis(cfg *configs.Config) (*RedisDB, error) {
	addr := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return &RedisDB{
		Redis: rdb,
	}, nil
}
