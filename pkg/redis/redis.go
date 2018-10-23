package redis

import (
	goredis "github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type client struct {
	redisClient *goredis.Client
}

type Client interface {
	Set(key string, value string) error
	Get(key string) (string, error)
}

func NewClient() (Client, error) {
	redisClient := goredis.NewClient(&goredis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		return client{}, errors.Wrap(err, "failed to connect to redis")
	}

	return client{
		redisClient: redisClient,
	}, nil
}

func (c client) Set(key string, value string) error {
	return c.redisClient.Set(key, value, 0).Err()
}

func (c client) Get(key string) (string, error) {
	return c.redisClient.Get(key).Result()
}
