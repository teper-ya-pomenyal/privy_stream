package domain

import "errors"

var (
	ErrTrackAlreadyExists = errors.New("track already exists")
	ErrTrackNotFound      = errors.New("track not found")
	ErrInvalidCharacters  = errors.New("invalid characters")
	ErrInvalidUUID        = errors.New("invalid uuid")
)
