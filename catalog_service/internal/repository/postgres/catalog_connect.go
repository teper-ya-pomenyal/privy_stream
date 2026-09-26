package postgres

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/config"
)

type PostgresCatalog struct {
	pool *pgxpool.Pool
}

func NewPostgresCatalog(cfg *config.Config) (*PostgresCatalog, error) {
	cfgPool, err := pgxpool.ParseConfig(cfg.DB.DSN)
	if err != nil {
		return nil, err
	}
	cfgPool.MaxConns = cfg.DB.MaxConns
	cfgPool.MinConns = cfg.DB.MinConns
	cfgPool.MaxConnLifetime = cfg.DB.MaxConnLifetime
	cfgPool.MaxConnIdleTime = cfg.DB.MaxConnIdleTime
	cfgPool.HealthCheckPeriod = cfg.DB.HealthCheckPeriod
	cfgPool.ConnConfig.ConnectTimeout = cfg.DB.ConnectTimeout
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	defer close()
	pool, err := pgxpool.NewWithConfig(ctx, cfgPool)
	if err != nil {
		return nil, err
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}
	return &PostgresCatalog{pool: pool}, nil
}

func (c *PostgresCatalog) Close() {
	c.pool.Close()
}
