package usecase

import (
	"context"

	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type RefreshResult struct {
	AccessToken  string
	RefreshToken string
}

type RefreshUseCase struct {
	tokenManager   TokenManager
	sessionManager SessionStore
}

func NewRefreshUseCase(tokenManager TokenManager, sessionManager SessionStore) *RefreshUseCase {
	return &RefreshUseCase{
		tokenManager:   tokenManager,
		sessionManager: sessionManager,
	}
}

func (r *RefreshUseCase) Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error) {
	userUUID, err := r.sessionManager.Get(ctx, refreshToken)
	if err == domain.ErrRefreshTokenNotFound {
		return nil, domain.ErrRefreshTokenNotFound
	}
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := r.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}

	newAccessToken, err := r.tokenManager.NewAccessToken(userUUID)
	if err != nil {
		return nil, err
	}

	result := &RefreshResult{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}
	err = r.sessionManager.Save(ctx, newRefreshToken, userUUID)
	if err != nil {
		return nil, err
	}
	err = r.sessionManager.Delete(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	return result, nil
}
