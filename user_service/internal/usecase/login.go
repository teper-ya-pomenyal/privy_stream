package usercase

import (
	"context"

	"github.com/teper-ya-pomenyal/privy_stream/jwtmanager"
	"github.com/teper-ya-pomenyal/privy_stream/user_service/internal/domain"
)

type UseUserCase struct {
	repo         domain.UsersRepository
	tokenManager jwtmanager.Manager
}

func (u *UseUserCase) Login(userName, password string) (string, string, string, string, error) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	user, err := u.repo.GetUserByUserName(ctx, userName)
	if err != nil {
		return "", "", "", "", err
	}

}
