package client

import (
	"context"
	"fmt"
	"ship/internal/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func newRedisClient(
	ctx context.Context,
	cfg *config.RedisConfig,
) (*RedisClient, error) {
	dbName, err := strconv.Atoi(cfg.DBName)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       dbName,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &RedisClient{client: client}, nil
}

func (c *RedisClient) Close() {
	c.client.Close()
}
