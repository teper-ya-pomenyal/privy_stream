package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/catalog_service/internal/usecase"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CatalogGRPCHandler struct {
	catalogv1.UnimplementedCatalogServiceServer
	trackUseCase usecase.TrackUseCase
}

func (h *CatalogGRPCHandler) GetTrackByID(ctx context.Context, req *catalogv1.GetTrackByIDRequest) (*catalogv1.GetTrackByIDResponse, error) {
	if req.TrackUuid == "" {
		return &catalogv1.GetTrackByIDResponse{}, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	trackUUID, err := uuid.Parse(req.TrackUuid)
	if err != nil {
		return &catalogv1.GetTrackByIDResponse{}, status.Error(codes.InvalidArgument, domain.ErrInvalidUUID.Error())
	}
	trackPath, err := h.trackUseCase.GetTrackForStream(ctx, trackUUID)
	if err != nil {
		return &catalogv1.GetTrackByIDResponse{}, mapDomainError(err)
	}
	return &catalogv1.GetTrackByIDResponse{
		Path:       trackPath.Path,
		DurationMs: int32(trackPath.DurationMS),
	}, nil
}
