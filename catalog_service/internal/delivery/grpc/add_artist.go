package grpc

import (
	"context"

	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CatalogWriteGRPCHandler) AddArtist(ctx context.Context, req *catalogv1.AddArtistRequest) (*catalogv1.AddArtistResponse, error) {
	if req.ArtistName == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artist, err := h.artistUseCase.AddArtist(ctx, req.ArtistName)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.AddArtistResponse{
		Artist: &catalogv1.Artist{
			ArtistUuid: artist.ArtistUUID.String(),
			ArtistName: artist.ArtistName,
		},
	}, nil
}
