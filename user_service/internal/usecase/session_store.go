package usecase

import (
	"context"

	"github.com/google/uuid"
)

type SessionStore interface {
	Save(ctx context.Context, refreshToken string, userUUID uuid.UUID) error
	Refresh(ctx context.Context, oldToken, newToken string) error
	Get(ctx context.Context, refreshToken string) (uuid.UUID, error)
	Delete(ctx context.Context, refreshToken string) error
}
