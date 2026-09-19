package grpc

import (
	"context"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *UserGRPCHandler) Logout(ctx context.Context, req *userv1.LogoutRequest) (*userv1.LogoutResponse, error) {
	if req.RefreshToken == "" {
		return &userv1.LogoutResponse{}, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}
	err := h.logoutUseCase.Logout(ctx, req.RefreshToken)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.LogoutResponse{}, nil
}
