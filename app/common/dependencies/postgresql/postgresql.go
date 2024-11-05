package postgresql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// DBConfig defines the interface that encapsulates database configuration information.
// It should be implemented by any type that provides a method to retrieve the connection string.
type DBConfig interface {
	ConnectionString() string
}

// NewPool builds a pool of pgx client.
func NewPool(cfg DBConfig) (PgxPoolIface, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, cfg.ConnectionString())
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return pool, nil
}
