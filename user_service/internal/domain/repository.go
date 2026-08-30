package domain

import "context"

type UsersRepository interface {
	AddUser(ctx context.Context, user *User) error
	GetUserByUserName(ctx context.Context, userName string) (*User, error)
	UserAlreadyExists(ctx context.Context, userName string) (bool, error)
}
