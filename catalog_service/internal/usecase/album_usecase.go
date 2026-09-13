package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

type AlbumUseCase struct {
	repo CatalogRepository
}

func NewAlbumUseCase(repo CatalogRepository) *AlbumUseCase {
	return &AlbumUseCase{repo: repo}
}

func (a *AlbumUseCase) GetAlbumByID(ctx context.Context, albumUUID uuid.UUID) (*domain.Album, error) {
	album, err := a.repo.GetAlbumByID(ctx, albumUUID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (a *AlbumUseCase) GetAlbumTracks(ctx context.Context, albumUUID uuid.UUID) ([]domain.LightAlbumTrack, error) {
	tracks, err := a.repo.GetAlbumTracks(ctx, albumUUID)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
