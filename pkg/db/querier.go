package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, params CreateUserParams) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
}

var _ Querier = (*Queries)(nil)
