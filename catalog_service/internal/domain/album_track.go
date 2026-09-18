package domain

import "github.com/google/uuid"

type AlbumTrack struct {
	TrackUUID uuid.UUID
	AlbumUUID uuid.UUID
	Position  int
}
