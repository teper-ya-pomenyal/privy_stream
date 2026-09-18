package domain

import "errors"

var (
	ErrTrackAlreadyExists  = errors.New("track already exists")
	ErrArtistAlreadyExists = errors.New("artist already exists")
	ErrAlbumAlreadyExists  = errors.New("album already exists")
	ErrTrackNotFound       = errors.New("track not found")
	ErrAlbumNotFound       = errors.New("album not found")
	ErrArtistNotFound      = errors.New("artist not found")
	ErrPlaylistNotFound    = errors.New("playlist not found")
	ErrInvalidCharacters   = errors.New("invalid characters")
	ErrInvalidUUID         = errors.New("invalid uuid")

	ErrAlbumTrackAlreadyExists   = errors.New("track is already added to this album")
	ErrAlbumTrackPositionTaken   = errors.New("position is already taken in this album")
	ErrAlbumTrackInvalidPosition = errors.New("track position must be greater than zero")
)

var (
	ErrInvalidPageParameters = errors.New("invalid page parameters")
	ErrInvalidTrackName      = errors.New("invalid track name")
	ErrInvalidArtistName     = errors.New("invalid artist name")
	ErrInvalidAlbumName      = errors.New("invalid album name")
)
