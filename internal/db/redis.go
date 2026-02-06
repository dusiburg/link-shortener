package db

import (
	"context"
	"log"
	"time"

	"example.com/m/internal/config"
	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis(cfg *config.Config) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.RedisURL)

	if err != nil {
		log.Fatalf("impossible to parse redis url: %v", err)
		return nil, err
	}

	Redis = redis.NewClient(opt)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := Redis.Ping(ctx).Err(); err != nil {
		log.Fatalf("impossible to conenct to redis db: %v", err)
		return nil, err
	}

	return Redis, nil
}
