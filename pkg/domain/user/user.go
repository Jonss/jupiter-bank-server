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

var (
	ErrUserExists      = errors.New("error user already exists")
	ErrPasswordInvalid = errors.New("error password is invalid")
	ErrUserIsDisabled  = errors.New("error user is disabled")
)

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

func (s *service) FetchUserByEmailAndPassword(ctx context.Context, email, password string) (*db.User, error) {
	user, err := s.queries.FetchUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = CheckPassword(user.PasswordHash, password)
	if err != nil {
		return nil, ErrPasswordInvalid
	}
	if !user.IsActive {
		return nil, ErrUserIsDisabled
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, userID uint64) (*db.User, error) {
	user, err := s.queries.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if !user.IsActive {
		return nil, ErrUserIsDisabled
	}
	return user, nil
}
