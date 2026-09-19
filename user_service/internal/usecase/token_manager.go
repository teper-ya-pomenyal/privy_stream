package usecase

import (
	"time"

	"github.com/google/uuid"
)

type TokenManager interface {
	NewAccessToken(userUUID uuid.UUID, birthDate time.Time) (string, error)
	NewRefreshToken() (string, error)
}
