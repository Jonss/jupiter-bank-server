package server

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/config"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/paseto_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/google/uuid"
	"time"
)

var fakeConfig = config.Config{
	Env: "test",
}

type userServiceMock struct {
	err      error
	wantUser *db.User
}

func (q *userServiceMock) FetchUserByEmailAndPassword(_ context.Context, email, password string) (*db.User, error) {
	if q.err != nil {
		return nil, q.err
	}
	q.wantUser.Email = email
	return q.wantUser, nil
}

func (q *userServiceMock) GetUserByID(_ context.Context, userID uint64) (*db.User, error) {
	if q.err != nil {
		return nil, q.err
	}
	q.wantUser.ID = userID
	return q.wantUser, nil
}

func (q *userServiceMock) CreateUser(_ context.Context, arg user.CreateUseParams) (*db.User, error) {
	if q.err != nil {
		return nil, q.err
	}
	externalID, _ := uuid.NewUUID()
	return &db.User{
		ExternalID:   externalID,
		Email:        arg.Email,
		Fullname:     arg.Fullname,
		PasswordHash: arg.Password,
		CreatedAt:    time.Now(),
	}, nil
}

type basicAuthMock struct {
	err error
}

func (d *basicAuthMock) FetchAppClient(_ context.Context, _ basic_auth.FetchAppClientParams) (bool, error) {
	if d.err != nil {
		return false, d.err
	}
	return true, nil
}

type pasetoAuthMock struct{}

func (m *pasetoAuthMock) Login(_ context.Context, email, password string) (paseto_auth.PasetoToken, error) {
	return paseto_auth.PasetoToken{}, nil
}

func (m *pasetoAuthMock) VerifyUser(_ context.Context, token, hex string) (*db.User, error) {
	return nil, nil
}
