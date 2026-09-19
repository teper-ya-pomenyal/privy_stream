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

func (h *CatalogGRPCHandler) GetTrackByID(ctx context.Context, req *catalogv1.GetTrackByIDRequest) (*catalogv1.GetTrackByIDResponse, error) {
	if req.TrackUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	trackUUID, err := uuid.Parse(req.TrackUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	trackPath, err := h.trackUseCase.GetTrackByID(ctx, trackUUID)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.GetTrackByIDResponse{
		Path:       trackPath.Path,
		DurationMs: int32(trackPath.DurationMS),
	}, nil
}

func (h *CatalogGRPCHandler) TrackExists(ctx context.Context, req *catalogv1.TrackExistsRequest) (*catalogv1.TrackExistsResponse, error) {
	if req.TrackUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	trackUUID, err := uuid.Parse(req.TrackUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	exists, err := h.trackUseCase.TrackExists(ctx, trackUUID)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.TrackExistsResponse{Exists: exists}, nil
}

func (h *CatalogGRPCHandler) SearchTrack(ctx context.Context, req *catalogv1.SearchTrackRequest) (*catalogv1.SearchTrackResponse, error) {
	tracks, err := h.trackUseCase.SearchTracks(ctx, req.TrackName, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, mapDomainError(err)
	}
	respTracks := make([]*catalogv1.Track, 0, len(tracks))
	for _, t := range tracks {
		respTracks = append(respTracks, &catalogv1.Track{
			TrackUuid:  t.TrackID.String(),
			TrackName:  t.TrackName,
			ArtistUuid: t.ArtistID.String(),
			ArtistName: t.ArtistName,
			AlbumUuid:  t.AlbumID.String(),
			AlbumName:  t.AlbumName,
			Explicit:   t.Explicit,
			DurationMs: int32(t.DurationMS),
		})
	}
	return &catalogv1.SearchTrackResponse{Tracks: respTracks}, nil
}

//////////////////////////////////////////

func (h *CatalogGRPCHandler) AddTrack(ctx context.Context, req *catalogv1.AddTrackRequest) (*catalogv1.AddTrackResponse, error) {
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
