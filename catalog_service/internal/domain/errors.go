package domain

import "errors"

var (
	ErrTrackAlreadyExists = errors.New("track already exists")
	ErrTrackNotFound      = errors.New("track not found")
	ErrAlbumNotFound      = errors.New("album not found")
	ErrArtistNotFound     = errors.New("artist not found")
	ErrPlaylistNotFound   = errors.New("playlist not found")
	ErrInvalidCharacters  = errors.New("invalid characters")
	ErrInvalidUUID        = errors.New("invalid uuid")
)

var (
	ErrInvalidPageParameters = errors.New("invalid page parameters")
	ErrInvalidTrackName      = errors.New("invalid track name")
	ErrInvalidArtistName     = errors.New("invalid artist name")
)
