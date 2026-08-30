package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserUUID     uuid.UUID
	UserName     string
	PasswordHash string
	BirthDate    time.Time
	CreatedAt    time.Time
}
