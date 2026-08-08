package client

import (
	"context"
	"fmt"
	"ship/internal/config"
)

type Client struct {
	Postgres *PostgresClient
	Redis    *RedisClient
	Resend   *ResendClient
}

func New(
	ctx context.Context,
	cfg *config.ClientConfig,
) (*Client, error) {
	pg, err := newPostgresClient(ctx, *cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres connection pool: %w", err)
	}

	rds, err := newRedisClient(ctx, cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("failed to create redis client: %w", err)
	}

	resend := newResendClient(cfg.Resend)

	return &Client{
		Postgres: pg,
		Redis:    rds,
		Resend:   resend,
	}, nil
}

func (c *Client) Close() {
	c.Postgres.Close()
	c.Redis.Close()
}
