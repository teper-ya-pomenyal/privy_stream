package clients

import (
	"context"
	"time"

	userv1 "github.com/teper-ya-pomenyal/privy_stream/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type LoginResult struct {
	UserUUID     string    `json:"user_uuid"`
	RefreshToken string    `json:"refresh_token"`
	AccessToken  string    `json:"access_token"`
	BirthDate    time.Time `json:"birth_date"`
}

type RefreshResult struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserClient struct {
	conn       *grpc.ClientConn
	grpcClient userv1.UserServiceClient
}

func NewUserClient(address string) (*UserClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &UserClient{conn: conn, grpcClient: userv1.NewUserServiceClient(conn)}, nil
}

func (u *UserClient) Login(ctx context.Context, userName, password string) (*LoginResult, error) {
	req := &userv1.LoginRequest{
		UserName: userName,
		Password: password,
	}
	res, err := u.grpcClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		UserUUID:     res.UserUuid,
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		BirthDate:    res.BirthDate.AsTime(),
	}, nil
}

func (u *UserClient) Register(ctx context.Context, userName, password string, birthDate time.Time) (*LoginResult, error) {
	req := &userv1.RegisterRequest{
		UserName:  userName,
		Password:  password,
		BirthDate: timestamppb.New(birthDate),
	}

	res, err := u.grpcClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return &LoginResult{
		UserUUID:     res.UserUuid,
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
		BirthDate:    res.BirthDate.AsTime(),
	}, nil
}

func (u *UserClient) Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error) {
	req := &userv1.RefreshRequest{RefreshToken: refreshToken}

	res, err := u.grpcClient.Refresh(ctx, req)
	if err != nil {
		return nil, err
	}
	return &RefreshResult{
		RefreshToken: res.RefreshToken,
		AccessToken:  res.AccessToken,
	}, nil
}

func (u *UserClient) Logout(ctx context.Context, refreshToken string) error {
	req := &userv1.LogoutRequest{RefreshToken: refreshToken}

	_, err := u.grpcClient.Logout(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
