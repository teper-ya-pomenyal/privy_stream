package grpc

import (
	"errors"

	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func mapDomainError(err error) error {
	switch {
	case errors.Is(err, domain.ErrTrackNotFound):
		return status.Error(codes.NotFound, err.Error())

	default:
		return status.Error(codes.Internal, "internal error")
	}
}
