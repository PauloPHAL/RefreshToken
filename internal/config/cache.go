package config

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
}

func NewCache() *Cache {
	config := GetConfig()

	return &Cache{
		client: redis.NewClient(&redis.Options{
			Addr: config.GetRedisHost() + ":" + strconv.Itoa(config.GetRedisPort()),
		}),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(context.Background(), key, data, ttl).Err()
}

func (c *Cache) Get(key string) ([]byte, error) {
	return c.client.Get(context.Background(), key).Bytes()
}
