package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type RedisRepository struct{}

func InitRedisRepository() *RedisRepository {
	return &RedisRepository{}
}

func (r *RedisRepository) StoreJWT(rc *redis.Client, key string, value interface{}) error {
	ctx := context.Background()
	valueBytes, _ := json.Marshal(value.(fiber.Map))
	stat := rc.Set(ctx, key, valueBytes, time.Duration(48)*time.Hour)
	return stat.Err()
}
