package db

import (
	"context"
	"database/sql"
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

func (u User) IsComplete() bool {
	return u.IsActive && u.TaxID != nil && u.Pin != nil
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

const fetchUserByEmailQuery = `
	SELECT 
		id, external_id, fullname,
		email, password_hash, pin, 
		tax_id, is_active,
		created_at, updated_at
	FROM users where email = $1;
`

func (q *Queries) FetchUserByEmail(ctx context.Context, email string) (*User, error) {
	row := q.db.QueryRowContext(ctx, fetchUserByEmailQuery, email)
	var u User
	err := row.Scan(
		&u.ID,
		&u.ExternalID,
		&u.Fullname,
		&u.Email,
		&u.PasswordHash,
		&u.Pin,
		&u.TaxID,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

const checkUserExistsByEmailQuery = `
	SELECT 
		id
	FROM users where email = $1;
`

func (q *Queries) CheckUserExistsByEmail(ctx context.Context, email string) (bool, error) {
	var ID int
	row := q.db.QueryRowContext(ctx, checkUserExistsByEmailQuery, email)
	err := row.Scan(
		&ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return ID > 0, nil
}

const fetchUserByIDQuery = `
	SELECT 
		id, external_id, fullname,
		email, password_hash, pin, 
		tax_id, is_active,
		created_at, updated_at
	FROM users where id = $1;
`

func (q *Queries) GetUserByID(ctx context.Context, ID uint64) (*User, error) {
	row := q.db.QueryRowContext(ctx, fetchUserByIDQuery, ID)
	var u User
	err := row.Scan(
		&u.ID,
		&u.ExternalID,
		&u.Fullname,
		&u.Email,
		&u.PasswordHash,
		&u.Pin,
		&u.TaxID,
		&u.IsActive,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
