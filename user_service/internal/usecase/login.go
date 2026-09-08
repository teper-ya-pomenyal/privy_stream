package usecase

import (
	"context"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/utils"
)

type LoginResult struct {
	UUID         string
	AccessToken  string
	RefreshToken string
	BirthDate    time.Time
	CreatedAt    time.Time
}

type LoginUseCase struct {
	repo           domain.UsersRepository
	tokenManager   TokenManager
	sessionManager SessionStore
}

func NewLoginUseCase(repo domain.UsersRepository, tokenManager TokenManager, sessionManager SessionStore) *LoginUseCase {
	return &LoginUseCase{
		repo:           repo,
		tokenManager:   tokenManager,
		sessionManager: sessionManager,
	}
}

func (l *LoginUseCase) Login(ctx context.Context, userName, password string) (*LoginResult, error) {
	// user search
	err := utils.ValidateUsername(userName)
	if err != nil {
		return &LoginResult{}, err
	}
	err = utils.ValidatePassword(password)
	if err != nil {
		return &LoginResult{}, err
	}

	user, err := l.repo.GetUserByUserName(ctx, userName)
	if err != nil {
		return &LoginResult{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return &LoginResult{}, domain.ErrInvalidCredentials
	}

	//make session

	accessToken, err := l.tokenManager.NewAccessToken(user.UserUUID)
	if err != nil {
		return &LoginResult{}, err
	}
	refreshToken, err := l.tokenManager.NewRefreshToken()
	if err != nil {
		return &LoginResult{}, err
	}
	loginResult := &LoginResult{
		UUID:         user.UserUUID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		BirthDate:    user.BirthDate,
		CreatedAt:    user.CreatedAt,
	}
	err = l.sessionManager.Save(ctx, refreshToken, user.UserUUID)
	if err != nil {
		return &LoginResult{}, err
	}
	return loginResult, nil
}
