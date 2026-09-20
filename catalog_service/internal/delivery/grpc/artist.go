package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *CatalogGRPCHandler) GetArtistByID(ctx context.Context, req *catalogv1.GetArtistByIDRequest) (*catalogv1.GetArtistByIDResponse, error) {
	if req.ArtistUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artistUUID, err := uuid.Parse(req.ArtistUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	artist, err := h.artistUseCase.GetArtistByID(ctx, artistUUID)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &catalogv1.GetArtistByIDResponse{
		Artist: &catalogv1.Artist{
			ArtistUuid: artist.ArtistUUID.String(),
			ArtistName: artist.ArtistName,
		},
	}, nil
}

func (h *CatalogGRPCHandler) SearchArtist(ctx context.Context, req *catalogv1.SearchArtistRequest) (*catalogv1.SearchArtistResponse, error) {
	artists, err := h.artistUseCase.SearchArtist(ctx, req.ArtistName, int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, mapDomainError(err)
	}
	respArtists := make([]*catalogv1.Artist, 0, len(artists))
	for _, a := range artists {
		respArtists = append(respArtists, &catalogv1.Artist{
			ArtistUuid: a.ArtistUUID.String(),
			ArtistName: a.ArtistName,
		})
	}
	return &catalogv1.SearchArtistResponse{Artists: respArtists}, nil
}

func (h *CatalogGRPCHandler) GetArtistAlbums(ctx context.Context, req *catalogv1.GetArtistAlbumsRequest) (*catalogv1.GetArtistAlbumsResponse, error) {
	if req.ArtistUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artistUUID, err := uuid.Parse(req.ArtistUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	albums, err := h.artistUseCase.GetArtistAlbums(ctx, artistUUID, req.Limit, req.Offset)
	if err != nil {
		return nil, mapDomainError(err)
	}
	respAlbums := make([]*catalogv1.LightAlbum, 0, len(albums))
	for _, a := range albums {
		respAlbums = append(respAlbums, &catalogv1.LightAlbum{
			AlbumUuid: a.AlbumUUID.String(),
			AlbumName: a.AlbumName,
			CreatedAt: a.CreatedAt.Format(timeLayout),
		})
	}
	return &catalogv1.GetArtistAlbumsResponse{ArtistAlbums: respAlbums}, nil
}

func (h *CatalogGRPCHandler) GetArtistTracks(ctx context.Context, req *catalogv1.GetArtistTracksRequest) (*catalogv1.GetArtistTracksResponse, error) {
	if req.ArtistUuid == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	artistUUID, err := uuid.Parse(req.ArtistUuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	tracks, err := h.artistUseCase.GetArtistTracks(ctx, artistUUID, req.Limit, req.Offset)
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
	return &catalogv1.GetArtistTracksResponse{ArtistTracks: respTracks}, nil
}
