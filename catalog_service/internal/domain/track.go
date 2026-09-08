package domain

import (
	"time"

	"github.com/google/uuid"
)

type Track struct {
	TrackID   uuid.UUID `db:"track_id"`
	TrackName string    `db:"track_name"`
	ArtistID  uuid.UUID `db:"artist_id"`
	AlbumID   uuid.UUID `db:"album_id"`
	Explicit  bool      `db:"explicit"`
	CreatedAt time.Time `db:"created_at"`
}
