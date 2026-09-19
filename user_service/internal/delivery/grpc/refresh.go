package grpc

import (
	"context"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *UserGRPCHandler) Refresh(ctx context.Context, req *userv1.RefreshRequest) (*userv1.RefreshResponse, error) {
	if req.RefreshToken == "" {
		return &userv1.RefreshResponse{}, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	res, err := h.refreshUseCase.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.RefreshResponse{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}, nil
}
