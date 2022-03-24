package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/google/uuid"
)

type CreateUseParams struct {
	Fullname string
	Email    string
	Password string
}

var ErrUserExists = errors.New("error user already exists")

func (d *service) CreateUser(ctx context.Context, arg CreateUseParams) (*db.User, error) {
	exists, err := d.queries.CheckUserExistsByEmail(ctx, arg.Email)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if exists {
		return nil, ErrUserExists
	}

	externalID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := HashPassword(arg.Password)
	if err != nil {
		return nil, err
	}

	user, err := d.queries.CreateUser(ctx, db.CreateUserParams{
		ExternalID:   externalID,
		Fullname:     arg.Fullname,
		Email:        arg.Email,
		PasswordHash: hashedPassword,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
