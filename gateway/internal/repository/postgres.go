package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DbPostgres struct {
	pool *pgxpool.Pool
}

func NewDbPostgres(ctx context.Context) (*DbPostgres, error) {
	pool, err := pgxpool.New(ctx, "postgres://marketplace:marketplace_dev@localhost:5432/auth?sslmode=disable")
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	return &DbPostgres{pool: pool}, nil
}

func (db *DbPostgres) Close() {
	db.pool.Close()
}
