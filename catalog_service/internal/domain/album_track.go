package domain

import "github.com/google/uuid"

type AlbumTrack struct {
	AlbumID  uuid.UUID `db:"album_id"`
	TrackID  uuid.UUID `db:"track_id"`
	Position int       `db:"position"`
}
