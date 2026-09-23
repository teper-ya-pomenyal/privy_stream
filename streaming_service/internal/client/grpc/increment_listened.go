package grpc

import (
	"context"

	"github.com/google/uuid"
	catalogv1 "github.com/teper-ya-pomenyal/privy_stream/proto/catalog/v1"
)

func (c *CatalogClient) IncrementListened(ctx context.Context, trackUUID uuid.UUID) error {
	_, err := c.writeClient.IncrementListened(ctx, &catalogv1.IncrementListenedRequest{TrackUuid: trackUUID.String()})
	return err
}
