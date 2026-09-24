package grpc

import (
	"context"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGRPCHandler) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	if req.UserName == "" {
		return nil, mapDomainError(domain.ErrLoginRequired)
	}

	if req.Password == "" {
		return nil, mapDomainError(domain.ErrPasswordRequired)
	}

	if req.BirthDate == nil {
		return nil, mapDomainError(domain.ErrBirthDateRequired)
	}

	res, err := h.registerUseCase.Register(ctx, req.UserName, req.Password, req.BirthDate.AsTime())
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.RegisterResponse{
		UserUuid:     res.UUID,
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		BirthDate:    timestamppb.New(res.BirthDate),
	}, nil
}
