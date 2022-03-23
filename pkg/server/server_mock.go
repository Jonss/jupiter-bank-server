package server

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/google/uuid"
	"time"
)

type userServiceMock struct {
	err error
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
