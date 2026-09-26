package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/config"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type UsersPostgresRepository struct {
	conn *pgxpool.Pool
}

func NewUsersPostgresRepository(cfg *config.UserDBConfig) (*UsersPostgresRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cfgDB, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, err
	}
	cfgDB.MaxConns = cfg.MaxConns
	cfgDB.MinConns = cfg.MinConns
	cfgDB.MaxConnLifetime = cfg.MaxConnLifetime
	cfgDB.MaxConnIdleTime = cfg.MaxConnIdleTime
	cfgDB.HealthCheckPeriod = cfg.HealthCheckPeriod
	cfgDB.ConnConfig.ConnectTimeout = cfg.ConnectTimeout
	conn, err := pgxpool.NewWithConfig(ctx, cfgDB)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(ctx); err != nil {
		conn.Close()
		return nil, err
	}

	return &UsersPostgresRepository{conn: conn}, nil
}

func (p *UsersPostgresRepository) Close() {
	p.conn.Close()
}

func (p *UsersPostgresRepository) AddUser(ctx context.Context, user *domain.User) error {
	_, err := p.conn.Exec(ctx,
		"INSERT INTO users (uuid, user_name, password_hash, birth_date, created_at) VALUES($1, $2, $3, $4, $5)",
		user.UserUUID, user.UserName, user.PasswordHash, user.BirthDate, user.CreatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.ErrUserAlreadyExists
		}
		return err
	}
	return nil

}

func (p *UsersPostgresRepository) GetUserByUserName(ctx context.Context, userName string) (*domain.User, error) {
	u := &domain.User{}
	err := p.conn.QueryRow(ctx,
		"SELECT uuid, user_name, password_hash, birth_date, created_at FROM users WHERE user_name = $1",
		userName,
	).Scan(&u.UserUUID, &u.UserName, &u.PasswordHash, &u.BirthDate, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func (p *UsersPostgresRepository) GetUserByID(ctx context.Context, userUUID uuid.UUID) (*domain.User, error) {
	u := &domain.User{}
	err := p.conn.QueryRow(ctx,
		"SELECT uuid, user_name, password_hash, birth_date, created_at FROM users WHERE uuid = $1",
		userUUID,
	).Scan(&u.UserUUID, &u.UserName, &u.PasswordHash, &u.BirthDate, &u.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}
	return u, nil
}

func (p *UsersPostgresRepository) UserAlreadyExists(ctx context.Context, userName string) (bool, error) {
	var exists bool
	err := p.conn.QueryRow(ctx,
		"SELECT EXISTS(SELECT 1 FROM users WHERE user_name = $1)",
		userName,
	).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil

}
