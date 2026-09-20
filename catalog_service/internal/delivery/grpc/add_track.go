package grpc

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CatalogWriteGRPCHandler) AddTrack(ctx context.Context, req *catalogv1.AddTrackRequest) (*catalogv1.AddTrackResponse, error) {
	if req.TrackName == "" || req.ArtistUuid == "" || req.AlbumUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artistUUID, err := uuid.Parse(req.ArtistUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	albumUUID, err := uuid.Parse(req.AlbumUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	track, err := h.trackUseCase.AddTrack(ctx, req.TrackName, artistUUID, albumUUID, req.Explicit, req.Path, time.Duration(req.DurationMs))
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.AddTrackResponse{
		TrackUuid:  track.TrackID.String(),
		TrackName:  track.TrackName,
		ArtistUuid: track.ArtistID.String(),
		AlbumUuid:  track.AlbumID.String(),
		Explicit:   track.Explicit,
		Path:       track.Path,
		DurationMs: int32(track.DurationMS),
	}, nil
}
