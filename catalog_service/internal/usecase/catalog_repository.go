package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

type CatalogRepository interface {
	GetTrackByID(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error)
	AddTrack(ctx context.Context, track *domain.Track) error
	TrackExists(ctx context.Context, trackUUID uuid.UUID) (bool, error)
	SearchTrack(ctx context.Context, trackName string, limit, offset int) ([]domain.Track, error)
	SearchArtist(ctx context.Context, artistName string, limit, offset int) ([]domain.Artist, error)
	GetArtistByID(ctx context.Context, artistUUID uuid.UUID) (*domain.Artist, error)
	GetArtistAlbums(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightAlbum, error)
	GetArtistTracks(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightTrack, error)
	GetAlbumByID(ctx context.Context, albumUUID uuid.UUID) (*domain.Album, error)
	GetAlbumTracks(ctx context.Context, albumUUID uuid.UUID) ([]domain.LightAlbumTrack, error)
}
