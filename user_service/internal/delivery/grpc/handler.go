package grpc

import (
	"context"
	"errors"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserGRPCHandler struct {
	userv1.UnimplementedUserServiceServer
	loginUseCase    *usecase.LoginUseCase
	registerUseCase *usecase.RegisterUseCase
	refreshUseCase  *usecase.RefreshUseCase
	logoutUseCase   *usecase.LogoutUseCase
}

func NewUserGRPCHandler(
	loginUseCase *usecase.LoginUseCase,
	registerUseCase *usecase.RegisterUseCase,
	refreshUseCase *usecase.RefreshUseCase,
	logoutUseCase *usecase.LogoutUseCase,
) *UserGRPCHandler {
	return &UserGRPCHandler{
		loginUseCase:    loginUseCase,
		registerUseCase: registerUseCase,
		refreshUseCase:  refreshUseCase,
		logoutUseCase:   logoutUseCase,
	}
}

func (h *UserGRPCHandler) Login(ctx context.Context, req *userv1.LoginRequest) (*userv1.LoginResponse, error) {
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

func (h *UserGRPCHandler) Register(ctx context.Context, req *userv1.RegisterRequest) (*userv1.RegisterResponse, error) {
	if req.BirthDate == nil {
		return nil, status.Error(codes.InvalidArgument, domain.ErrInvalidDate.Error())
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

func (h *UserGRPCHandler) Refresh(ctx context.Context, req *userv1.RefreshRequest) (*userv1.RefreshResponse, error) {
	res, err := h.refreshUseCase.Refresh(ctx, req.RefreshToken)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.RefreshResponse{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}, nil
}

func (h *UserGRPCHandler) Logout(ctx context.Context, req *userv1.LogoutRequest) (*userv1.LogoutResponse, error) {
	err := h.logoutUseCase.Logout(ctx, req.RefreshToken)
	if err != nil {
		return nil, mapDomainError(err)
	}
	return &userv1.LogoutResponse{}, nil
}

// errors mapping
func mapDomainError(err error) error {
	switch {
	case errors.Is(err, domain.ErrInvalidCredentials):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, domain.ErrUserNotFound):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, domain.ErrUserAlreadyExists):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, domain.ErrRefreshTokenNotFound):
		return status.Error(codes.Unauthenticated, err.Error())
	default:
		return status.Error(codes.Internal, "internal error")
	}
}
