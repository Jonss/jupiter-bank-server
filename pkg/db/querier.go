package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, params CreateUserParams) (*User, error)
	FetchUserByEmail(ctx context.Context, email string) (*User, error)
	CheckUserExistsByEmail(ctx context.Context, email string) (bool, error)
	FetchAppClient(ctx context.Context, arg FetchAppClientParams) (*AppClient, error)
	GetUserByID(ctx context.Context, ID uint64) (*User, error)
}

var _ Querier = (*Queries)(nil)
