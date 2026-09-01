package usecase

import "github.com/google/uuid"

type TokenManager interface {
	NewAccessToken(userUUID uuid.UUID) (string, error)
	NewRefreshToken() (string, error)
}
