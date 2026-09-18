package grpc

import (
	"errors"
	"time"

	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const timeLayout = time.RFC3339

func mapDomainError(err error) error {
	switch {
	case errors.Is(err, domain.ErrTrackNotFound),
		errors.Is(err, domain.ErrArtistNotFound),
		errors.Is(err, domain.ErrAlbumNotFound),
		errors.Is(err, domain.ErrPlaylistNotFound):
		return status.Error(codes.NotFound, err.Error())

	case errors.Is(err, domain.ErrTrackAlreadyExists),
		errors.Is(err, domain.ErrArtistAlreadyExists),
		errors.Is(err, domain.ErrAlbumAlreadyExists),
		errors.Is(err, domain.ErrAlbumTrackAlreadyExists),
		errors.Is(err, domain.ErrAlbumTrackPositionTaken):
		return status.Error(codes.AlreadyExists, err.Error())

	case errors.Is(err, domain.ErrInvalidPageParameters),
		errors.Is(err, domain.ErrInvalidTrackName),
		errors.Is(err, domain.ErrInvalidArtistName),
		errors.Is(err, domain.ErrInvalidAlbumName),
		errors.Is(err, domain.ErrAlbumTrackInvalidPosition),
		errors.Is(err, domain.ErrInvalidCharacters),
		errors.Is(err, domain.ErrInvalidUUID):
		return status.Error(codes.InvalidArgument, err.Error())

	default:
		return status.Error(codes.Internal, "internal error")
	}
}
