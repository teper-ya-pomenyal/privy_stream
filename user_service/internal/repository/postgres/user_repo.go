package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type UsersPostgresRepository struct {
	conn *sqlx.DB
}

func NewUsersPostgresRepository(dsn string) (*UsersPostgresRepository, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}
	return &UsersPostgresRepository{conn: db}, nil
}

func (p *UsersPostgresRepository) AddUser(ctx context.Context, user *domain.User) error {
	_, err := p.conn.ExecContext(ctx,
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
	user := &domain.User{}
	err := p.conn.GetContext(ctx, user,
		"SELECT uuid, user_name, password_hash, birth_date, created_at FROM users WHERE user_name = $1",
		userName,
	)
	switch err {
	case sql.ErrNoRows:
		return nil, domain.ErrUserNotFound
	case nil:
		return user, nil
	default:
		return nil, err
	}

}

func (p *UsersPostgresRepository) UserAlreadyExists(ctx context.Context, userName string) (bool, error) {
	var exists bool
	err := p.conn.GetContext(ctx, &exists,
		"SELECT EXISTS(SELECT 1 FROM users WHERE user_name = $1)",
		userName,
	)
	if err != nil {
		return false, err
	}
	return exists, nil

}
