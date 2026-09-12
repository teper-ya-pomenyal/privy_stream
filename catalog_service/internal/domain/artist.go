package domain

import (
	"github.com/google/uuid"
)

type Artist struct {
	ArtistID   uuid.UUID `db:"artist_id"`
	ArtistName string    `db:"artist_name"`
}
