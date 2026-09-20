package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CatalogWriteGRPCHandler) AddTracksToAlbum(ctx context.Context, req *catalogv1.AddTracksToAlbumRequest) (*catalogv1.AddTracksToAlbumResponse, error) {
	if req.AlbumUuid == "" || len(req.Tracks) == 0 {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	albumUUID, err := uuid.Parse(req.AlbumUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	albumTracks := make([]domain.AlbumTrack, 0, len(req.Tracks))
	for _, t := range req.Tracks {
		trackUUID, err := uuid.Parse(t.TrackUuid)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
		}
		albumTracks = append(albumTracks, domain.AlbumTrack{
			TrackUUID: trackUUID,
			AlbumUUID: albumUUID,
			Position:  int(t.Position),
		})
	}
	if err := h.albumUseCase.AddTracksToAlbum(ctx, albumTracks); err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.AddTracksToAlbumResponse{}, nil
}
