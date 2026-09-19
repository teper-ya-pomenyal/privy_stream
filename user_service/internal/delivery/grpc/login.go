package grpc

import (
	"context"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *UserGRPCHandler) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {

	if req.UserName == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}

	if req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidCharacters.Error())
	}

	res, err := h.loginUseCase.Login(ctx, req.UserName, req.Password)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.LoginResponse{
		UserUuid:     res.UUID,
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		BirthDate:    timestamppb.New(res.BirthDate),
	}, nil
}
