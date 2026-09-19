package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CatalogGRPCHandler) GetAlbumByID(ctx context.Context, req *catalogv1.GetAlbumByIDRequest) (*catalogv1.GetAlbumByIDResponse, error) {
	if req.AlbumUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	albumUUID, err := uuid.Parse(req.AlbumUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	album, err := h.albumUseCase.GetAlbumByID(ctx, albumUUID)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.GetAlbumByIDResponse{
		Album: &catalogv1.Album{
			AlbumUuid:  album.AlbumUUID.String(),
			ArtistUuid: album.ArtistUUID.String(),
			AlbumName:  album.AlbumName,
			CreatedAt:  album.CreatedAt.Format(timeLayout),
		},
	}, nil
}

func (h *CatalogGRPCHandler) GetAlbumTracks(ctx context.Context, req *catalogv1.GetAlbumTracksRequest) (*catalogv1.GetAlbumTracksResponse, error) {
	if req.AlbumUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	albumUUID, err := uuid.Parse(req.AlbumUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	tracks, err := h.albumUseCase.GetAlbumTracks(ctx, albumUUID)
	if err != nil {
		return nil, mapDomainError(err)
	}
	respTracks := make([]*catalogv1.LightTrack, 0, len(tracks))
	for _, t := range tracks {
		respTracks = append(respTracks, &catalogv1.LightTrack{
			TrackUuid:  t.TrackID.String(),
			TrackName:  t.TrackName,
			Explicit:   t.Explicit,
			DurationMs: int32(t.DurationMS),
		})
	}
	return &catalogv1.GetAlbumTracksResponse{LightTrack: respTracks}, nil
}

//////////////////////////////////////////

func (h *CatalogGRPCHandler) AddAlbum(ctx context.Context, req *catalogv1.AddAlbumRequest) (*catalogv1.AddAlbumResponse, error) {
	if req.ArtistUuid == "" || req.AlbumName == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artistUUID, err := uuid.Parse(req.ArtistUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	album, err := h.albumUseCase.AddAlbum(ctx, artistUUID, req.AlbumName)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.AddAlbumResponse{
		Album: &catalogv1.Album{
			AlbumUuid:  album.AlbumUUID.String(),
			ArtistUuid: album.ArtistUUID.String(),
			AlbumName:  album.AlbumName,
			CreatedAt:  album.CreatedAt.Format(timeLayout),
		},
	}, nil
}

func (h *CatalogGRPCHandler) AddTracksToAlbum(ctx context.Context, req *catalogv1.AddTracksToAlbumRequest) (*catalogv1.AddTracksToAlbumResponse, error) {
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
