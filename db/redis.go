package db

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func StartRedisConnection() *redis.Client {
	username := viper.GetString("REDIS_USERNAME")
	password := viper.GetString("REDIS_PASSWORD")
	host := viper.GetString("REDIS_HOST")
	port := viper.GetString("REDIS_PORT")

	client := redis.NewClient(&redis.Options{
		Username:     username,
		Password:     password,
		Addr:         fmt.Sprintf("%s:%s", host, port),
		PoolTimeout:  10 * time.Minute,
		MinIdleConns: 10,
	})
	return client
}
