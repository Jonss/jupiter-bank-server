package paseto_auth

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
)

type service struct {
	userService user.Service
}

func NewPasetoAuthService(userService user.Service) service {
	return service{userService: userService}
}

type Service interface {
	Login(ctx context.Context, email, password string) (PasetoToken, error)
	VerifyUser(ctx context.Context, token, hex string) (*db.User, error)
}
