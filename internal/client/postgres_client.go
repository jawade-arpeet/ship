package client

import (
	"context"
	"errors"
	"fmt"
	"ship/internal/config"
	"ship/internal/errs"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var errPgOperationFailed = errors.New("postgres operation failed")

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

func (c *PostgresClient) wrapErr(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return errs.ErrPgNoRows
	}

	pgErr, ok := errors.AsType[*pgconn.PgError](err)
	if !ok {
		return fmt.Errorf("%w: %w", errPgOperationFailed, err)
	}

	var sentinel error
	switch pgErr.Code {
	case "23505":
		sentinel = errs.ErrPgUniqueViolation
	case "23503":
		sentinel = errs.ErrPgForeignKeyViolation
	case "23502":
		sentinel = errs.ErrPgNotNullViolation
	case "23514":
		sentinel = errs.ErrPgCheckViolation
	default:
		return fmt.Errorf("%w: %w", errPgOperationFailed, err)
	}

	return fmt.Errorf("%w: %w: %w", errPgOperationFailed, sentinel, pgErr)
}

func (c *PostgresClient) QueryRow(
	ctx context.Context,
	query string,
	args pgx.NamedArgs,
	result any,
) error {
	rows, err := c.pool.Query(ctx, query, args)
	if err != nil {
		return c.wrapErr(err)
	}

	defer rows.Close()

	if err := pgxscan.ScanOne(result, rows); err != nil {
		return c.wrapErr(err)
	}

	return nil
}

func (c *PostgresClient) Query(
	ctx context.Context,
	query string,
	args pgx.NamedArgs,
	result any,
) error {
	rows, err := c.pool.Query(ctx, query, args)
	if err != nil {
		return c.wrapErr(err)
	}

	defer rows.Close()

	if err := pgxscan.ScanAll(result, rows); err != nil {
		return c.wrapErr(err)
	}

	return nil
}
