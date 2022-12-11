package auth

import (
	"github.com/go-redis/redis/v9"
	"os"
)

type RedisService struct {
	Auth   IAuthentication
	Client *redis.Client
}

var (
	host     = os.Getenv("REDIS_HOST")
	port     = os.Getenv("REDIS_PORT")
	password = os.Getenv("REDIS_PASSWORD")
)

func NewRedisDB() (*RedisService, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	return &RedisService{
		Auth:   NewAuth(redisClient),
		Client: redisClient,
	}, nil
}
