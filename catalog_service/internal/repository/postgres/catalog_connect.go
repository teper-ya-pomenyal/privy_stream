package postgres

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/config"
)

type PostgresCatalog struct {
	conn *sqlx.DB
}

func NewPostgresCatalog(cfg *config.Config) (*PostgresCatalog, error) {
	db, err := sqlx.Connect("pgx", cfg.DSN)
	if err != nil {
		return nil, err
	}
	return &PostgresCatalog{conn: db}, nil
}
