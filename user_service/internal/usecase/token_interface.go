package usecase

type TokenManager interface {
	NewAccessToken(userUUID string) (string, error)
	NewRefreshToken() (string, error)
}
