package repositories

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type LinksRepository struct {
	redis *redis.Client
}

func NewLinksRepository(redis *redis.Client) *LinksRepository {
	return &LinksRepository{redis: redis}
}

func (r *LinksRepository) CreateLink(shortID, original string) error {
	return r.redis.Set(context.Background(), shortID, original, 30*24*time.Hour).Err()
}

func (r *LinksRepository) GetLink(shortID string) (string, error) {
	return r.redis.Get(context.Background(), shortID).Result()
}
