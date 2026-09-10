package domain

import (
	"time"

	"github.com/google/uuid"
)

type Artist struct {
	ArtistID   uuid.UUID `db:"artist_id"`
	ArtistName string    `db:"artist_name"`
	Tracks     []Track
	Albums     []Album
	CreatedAt  time.Time `db:"created_at"`
}

type LightArtist struct {
	ArtistID   uuid.UUID `db:"artist_id"`
	ArtistName string    `db:"artist_name"`
}
