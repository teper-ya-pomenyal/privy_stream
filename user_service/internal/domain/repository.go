package domain

import (
	"context"

	"github.com/google/uuid"
)

type UsersRepository interface {
	AddUser(ctx context.Context, user *User) error
	GetUserByUserName(ctx context.Context, userName string) (*User, error)
	GetUserByID(ctx context.Context, userUUID uuid.UUID) (*User, error)
	UserAlreadyExists(ctx context.Context, userName string) (bool, error)
}
