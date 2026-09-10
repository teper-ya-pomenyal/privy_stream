package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

type TrackUseCase struct {
	repo CatalogRepository
}

func (t *TrackUseCase) GetTrackForStream(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error) {
	trackPath, err := t.repo.GetTrackByID(ctx, trackUUID)
	if err != nil {
		return &domain.TrackPath{}, err
	}
	return trackPath, nil
}

//func (t *TrackUseCase)
