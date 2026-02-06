package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerHost string
	ServerPort string
	RedisURL   string
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		ServerHost: os.Getenv("SERVER_HOST"),
		ServerPort: os.Getenv("SERVER_PORT"),
		RedisURL:   os.Getenv("REDIS_URL"),
	}

	if cfg.ServerHost == "" {
		cfg.ServerHost = "localhost"
	}

	if cfg.ServerPort == "" {
		return nil, fmt.Errorf("server port is required")
	}

	if cfg.RedisURL == "" {
		return nil, fmt.Errorf("redis connection url is required")
	}

	return cfg, nil
}
