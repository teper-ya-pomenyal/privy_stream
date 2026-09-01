package usecase

import "context"

type LogoutUseCase struct {
	sessionManager SessionStore
}

func NewLogoutUseCase(sessionManager SessionStore) *LogoutUseCase {
	return &LogoutUseCase{
		sessionManager: sessionManager,
	}
}

func (l *LogoutUseCase) Logout(ctx context.Context, refreshToken string) error {
	err := l.sessionManager.Delete(ctx, refreshToken)
	if err != nil {
		return err
	}
	return nil
}
