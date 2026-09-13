package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/utilites"
)

type TrackUseCase struct {
	repo CatalogRepository
}

func NewTrackUserCase(repo CatalogRepository) *TrackUseCase {
	return &TrackUseCase{repo: repo}
}

func (t *TrackUseCase) GetTrackByID(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error) {
	trackPath, err := t.repo.GetTrackByID(ctx, trackUUID)
	if err != nil {
		return nil, err
	}
	return trackPath, nil
}

func (t *TrackUseCase) TrackExists(ctx context.Context, trackUUID uuid.UUID) (bool, error) {
	ok, err := t.repo.TrackExists(ctx, trackUUID)
	return ok, err
}

func (t *TrackUseCase) SearchTracks(ctx context.Context, trackName string, limit, offset int) ([]domain.Track, error) {
	if limit < 1 || offset < 0 {
		return nil, domain.ErrInvalidPageParameters
	}
	cleanTN, err := utilites.ValidateTrackName(trackName)
	if err != nil {
		return nil, err
	}
	tracks, err := t.repo.SearchTrack(ctx, cleanTN, limit, offset)
	if err != nil {
		return nil, err
	}
	return tracks, nil
}
