package domain

import "github.com/google/uuid"

type PlaylistTrack struct {
	PlaylistUUID uuid.UUID `db:"playlist_id"`
	TrackID      uuid.UUID `db:"track_id"`
	Position     int       `db:"position"`
}
