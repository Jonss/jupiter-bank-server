package user

import (
	"context"
	"fmt"

	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/google/uuid"
)

func (d *UserDomain) createUser(ctx context.Context, req createUserRequest) (*db.User, error) {
	externalID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	fmt.Println("UserRequest:", req)
	user, err := d.q.CreateUser(ctx, db.CreateUserParams{
		ExternalID:   externalID,
		Fullname:     req.Fullname,
		Email:        req.Email,
		PasswordHash: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
