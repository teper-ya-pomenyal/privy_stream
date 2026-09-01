package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserUUID     uuid.UUID `db:"uuid"`
	UserName     string    `db:"user_name"`
	PasswordHash string    `db:"password_hash"`
	BirthDate    time.Time `db:"birth_date"`
	CreatedAt    time.Time `db:"created_at"`
}
