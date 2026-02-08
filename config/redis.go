package config

import (
	"exchange_backend/global"
	"github.com/go-redis/redis"
	"log"
)

func initRedis() {
	RedisClient := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis,got error: %v", err)
	}
	global.RedisDB = RedisClient
}
