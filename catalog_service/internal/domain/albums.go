package domain

import (
	"time"

	"github.com/google/uuid"
)

type Album struct {
	AlbumID   uuid.UUID `db:"album_id"`
	ArtistID  uuid.UUID `db:"artist_id"`
	AlbumName string    `db:"artist_name"`
	CreatedAt time.Time `db:"created_as"`
}
