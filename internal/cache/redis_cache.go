package cache

import (
	"github.com/go-redis/redis"
	"github.com/ngonghi/admin_site/config"
	"log"
	"time"
)

type RedisCache struct {
	Client *redis.Client
}

func (c *RedisCache) Ping() error {
	return c.Client.Ping().Err()
}

func (c *RedisCache) Get(key string) (string, error) {
	return c.Client.Get(key).Result()
}

func (c *RedisCache) Set(key string, value interface{}, time time.Duration) (string, error) {
	return c.Client.Set(key, value, time).Result()
}

func NewRedisClient(c *config.Configuration) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     c.RedisAddr,
		Password: c.RedisPwd,
		DB:       0, // use default DB
	})

	pong, err := client.Ping().Result()

	if err != nil || pong == "" {
		log.Fatalf("redis cache: got no PONG back %q", err)
	}

	return client
}
