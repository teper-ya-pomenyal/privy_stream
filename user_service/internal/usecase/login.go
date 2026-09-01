package usecase

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type LoginResult struct {
	UUID         string
	AccessToken  string
	RefreshToken string
	BirthDate    time.Time
	CreatedAt    time.Time
}

type LoginUseCase struct {
	repo         domain.UsersRepository
	tokenManager TokenManager
}

func (l *LoginUseCase) Login(ctx context.Context, userName, password string) (*LoginResult, error) {
	user, err := l.repo.GetUserByUserName(ctx, userName)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, domain.ErrInvalidCredentials
	}
	accessToken, err := l.tokenManager.NewAccessToken(user.UserUUID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := l.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		UUID:         user.UserUUID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		BirthDate:    user.BirthDate,
		CreatedAt:    user.CreatedAt,
	}, nil
}
