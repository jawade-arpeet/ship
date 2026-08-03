package client

import (
	"context"
	"fmt"
	"ship/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(
	ctx context.Context,
	cfg config.PostgresConfig,
) (*PostgresClient, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres client: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres client: %w", err)
	}

	return &PostgresClient{pool: pool}, nil
}

func (c *PostgresClient) Close() {
	c.pool.Close()
}
