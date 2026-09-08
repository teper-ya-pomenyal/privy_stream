package domain

import (
	"time"

	"github.com/google/uuid"
)

type Playlist struct {
	PlaylistID   uuid.UUID `db:"playlist_id"`
	PlaylistName string    `db:"playlist_name"`
	OwnerID      uuid.UUID `db:"owner_id"`
	CreatedAt    time.Time `db:"created_at"`
}
