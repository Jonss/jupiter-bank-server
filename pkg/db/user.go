package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// model
type User struct {
	ID           uint64
	ExternalID   uuid.UUID
	Fullname     string
	Email        string
	PasswordHash string
	Pin          *string
	TaxID        *string
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateUserParams struct {
	ID           uint64
	ExternalID   uuid.UUID
	Fullname     string
	Email        string
	PasswordHash string
}

const createUserQuery = `
	INSERT INTO users(
		external_id,
		fullname,
		email,
		password_hash
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) RETURNING id, external_id, fullname, email, created_at, updated_at;
`

func (q *Queries) CreateUser(ctx context.Context, params CreateUserParams) (*User, error) {
	row := q.db.QueryRowContext(ctx, createUserQuery, params.ExternalID, params.Fullname, params.Email, params.PasswordHash)
	var u User
	err := row.Scan(
		&u.ID,
		&u.ExternalID,
		&u.Fullname,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	return &u, err
}
