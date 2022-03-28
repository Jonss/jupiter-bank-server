package user

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
)

type service struct {
	queries db.Querier
}

func NewUserService(q db.Querier) *service {
	return &service{queries: q}
}

type Service interface {
	CreateUser(ctx context.Context, arg CreateUseParams) (*db.User, error)
	FetchUserByEmailAndPassword(ctx context.Context, email, password string) (*db.User, error)
	GetUserByID(ctx context.Context, userID uint64)(*db.User, error)
}

var _ Service = (*service)(nil)
