package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUseCase struct {
	repo         domain.UsersRepository
	tokenManager TokenManager
}

func (r *RegisterUseCase) Register(ctx context.Context, userName, password, birthDate string) (*LoginResult, error) {

	ok, err := r.repo.UserAlreadyExists(ctx, userName)
	if err != nil {
		return nil, err
	}
	if ok {
		return nil, domain.ErrUserAlreadyExists
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10) //установил дефолтное значение для ясности - какое колличество рацндов хэширования.
	if err != nil {
		return nil, err
	}
	t, err := time.Parse(time.DateOnly, birthDate)
	if err != nil {
		return nil, err
	}
	newUser := &domain.User{
		UserUUID:     uuid.New(),
		UserName:     userName,
		PasswordHash: string(hashPassword),
		BirthDate:    t,
		CreatedAt:    time.Now(),
	}

	err = r.repo.AddUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	accessToken, err := r.tokenManager.NewAccessToken(newUser.UserUUID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := r.tokenManager.NewRefreshToken()
	if err != nil {
		return nil, err
	}
	resultLogin := &LoginResult{
		UUID:         newUser.UserUUID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		BirthDate:    newUser.BirthDate,
		CreatedAt:    newUser.CreatedAt,
	}
	return resultLogin, nil
}
