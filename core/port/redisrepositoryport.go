package port

import "github.com/go-redis/redis/v8"

type RedisRepositoryInterface interface {
	StoreJWT(rc *redis.Client, key string, value interface{}) error
}
