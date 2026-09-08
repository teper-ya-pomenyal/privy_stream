package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
)

type CatalogRepository interface {
	GetTrackById(ctx context.Context, trackUUID uuid.UUID) (*domain.TrackPath, error)
}
