package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/utilites"
)

type ArtistUseCase struct {
	repo CatalogRepository
}

func NewArtistRepository(repo CatalogRepository) *ArtistUseCase {
	return &ArtistUseCase{repo: repo}
}

func (a *ArtistUseCase) SearchArtist(ctx context.Context, artistName string, limit, offset int) ([]domain.Artist, error) {
	if limit < 1 || offset < 0 {
		return nil, domain.ErrInvalidPageParameters
	}
	cleanAN, err := utilites.ValidateArtistName(artistName)
	if err != nil {
		return nil, err
	}

	artists, err := a.repo.SearchArtist(ctx, cleanAN, limit, offset)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func (a *ArtistUseCase) GetArtistByID(ctx context.Context, artistUUID uuid.UUID) (*domain.Artist, error) {
	artist, err := a.repo.GetArtistByID(ctx, artistUUID)
	if err != nil {
		return nil, err
	}
	return artist, nil
}

func (a *ArtistUseCase) GetArtistAlbums(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.Album, error) {
	if limit < 1 || offset < 0 {
		return nil, domain.ErrInvalidPageParameters
	}
	albums, err := a.repo.GetArtistAlbums(ctx, artistUUID, limit, int32(offset))
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (a *ArtistUseCase) GetArtistTracks(ctx context.Context, artistUUID uuid.UUID, limit, offset int32) ([]domain.LightTrack, error) {
	if limit < 1 || offset < 0 {
		return nil, domain.ErrInvalidPageParameters
	}
	tracks, err := a.repo.GetArtistTracks(ctx, artistUUID, limit, offset)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}

func (a *ArtistUseCase) AddArtist(ctx context.Context, artistName string) (*domain.Artist, error) {
	cleanAN, err := utilites.ValidateArtistName(artistName)
	if err != nil {
		return nil, err
	}

	artist := domain.Artist{
		ArtistUUID: uuid.New(),
		ArtistName: cleanAN,
		CreatedAt:  time.Now(),
	}
	if err := a.repo.AddArtist(ctx, artist); err != nil {
		return nil, err
	}
	return &artist, nil
}
