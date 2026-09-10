package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

type CatalogRepository interface {
	GetTrackByID(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error)
	AddTrack(ctx context.Context, track *domain.Track) error
	GetTracksPage(ctx context.Context, trackName string, limit, offset int) ([]domain.Track, error)
	SearchArtist(ctx context.Context, artistName string, limit, offset int) ([]domain.LightArtist, error)
	GetArtistAlbums(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightAlbum, error)
	GetArtistTracks(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightTrack, error)
}
