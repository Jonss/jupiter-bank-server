package basic_auth

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
)

type service struct {
	queries db.Querier
}

func NewBasicAuthService(q db.Querier) *service {
	return &service{queries: q}
}

type Service interface {
	FetchAppClient(ctx context.Context, arg FetchAppClientParams) (bool, error)
}

var _ Service = (*service)(nil)
