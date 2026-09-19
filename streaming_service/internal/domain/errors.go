package domain

import "errors"

var (
	ErrFileNotExists          = errors.New("track file not exists")
	ErrNotPermission          = errors.New("not permission")
	ErrTrackNotFound          = errors.New("track not found")
	ErrInternalServer         = errors.New("internal server error")
	ErrExplicitContentBlocked = errors.New("explicit content blocked")
)
