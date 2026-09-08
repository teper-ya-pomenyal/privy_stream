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
	// validation refresh token
	userUUID, err := r.sessionManager.Get(ctx, refreshToken)
	if err == domain.ErrRefreshTokenNotFound {
		return nil, domain.ErrRefreshTokenNotFound
	}
	if err != nil {
		return &RefreshResult{}, err
	}

	// make response
	newRefreshToken, err := r.tokenManager.NewRefreshToken()
	if err != nil {
		return &RefreshResult{}, err
	}

	newAccessToken, err := r.tokenManager.NewAccessToken(userUUID)
	if err != nil {
		return &RefreshResult{}, err
	}

	result := &RefreshResult{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}
	err = r.sessionManager.Refresh(ctx, refreshToken, newRefreshToken)
	if err != nil {
		return &RefreshResult{}, err
	}
	return result, nil
}
