package domain

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ArtistUUID uuid.UUID `db:"artist_id"`
	ArtistName string    `db:"artist_name"`
	CreatedAt  time.Time `db:"created_at"`
}
