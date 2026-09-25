package client

import (
	"context"
	"fmt"
	"ship/internal/config"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClient struct {
	pool *pgxpool.Pool
}

func newPostgresClient(ctx context.Context) (*PostgresClient, error) {
	cfg := config.GetPostgresConfig()

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres client: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresClient{pool: pool}, nil
}

func (c *PostgresClient) Close() {
	c.pool.Close()
}

func (c *PostgresClient) QueryOne[T any](
	ctx context.Context,
	query string,
	args pgx.NamedArgs,
) (*T, error) {
	row, err := c.pool.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var result T
	if err := pgxscan.ScanOne(&result, row); err != nil {
		return nil, err
	}

	return &result, nil
}
