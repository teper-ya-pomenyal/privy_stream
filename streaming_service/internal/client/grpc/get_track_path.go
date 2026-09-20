package grpc

import (
	"context"

	"github.com/google/uuid"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
	"github.com/teper-ya-pomenyal/privy_stream/streaming_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *CatalogClient) GetTrackByID(ctx context.Context, trackUUID uuid.UUID) (*TrackLocation, error) {
	res, err := c.client.GetTrackByID(ctx, &catalogv1.GetTrackByIDRequest{TrackUuid: trackUUID.String()})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, err
		}
		switch st.Code() {
		case codes.NotFound:
			return nil, domain.ErrTrackNotFound
		default:
			return nil, err
		}
	}
	return &TrackLocation{Path: res.Path, DurationMS: res.DurationMs, Explicit: res.Explicit}, nil
}
