package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func createUser(t *testing.T) {
	insertQuery := `
	INSERT INTO users(
		external_id,
		fullname,
		email,
		password_hash
	) VALUES (
		'74b2a86e-a975-11ec-b909-0242ac120002',
		'Existing user',
		'existing.user@email.com',
		'a_hashed_password'
	)`

	_, err := testQueries.db.ExecContext(context.Background(), insertQuery)
	if err != nil {
		t.Fatalf("unexpected error inserting config. error=(%v)", err)
	}
}

func TestCreateUser(t *testing.T) {
	createUser(t)

	uuid1, _ := uuid.NewUUID()
	uuid2, _ := uuid.NewUUID()
	testCases := []struct {
		name         string
		arg          CreateUserParams
		hasError     bool
		errorCode    string
		errorMessage string
	}{
		{
			name: "should create user successfully",
			arg: CreateUserParams{
				Fullname:   "Jupiter Stein",
				Email:      "jupiter.stein@jupiterbank.co",
				ExternalID: uuid1,
			},
		},
		{
			name: "should not create user email already exists",
			arg: CreateUserParams{
				Fullname:   "Franz Ferdinand",
				Email:      "existing.user@email.com",
				ExternalID: uuid2,
			},
			hasError:     true,
			errorCode:    "23505",
			errorMessage: `duplicate key value violates unique constraint "users_email_key"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := testQueries.CreateUser(context.Background(), tc.arg)
			if tc.hasError {
				if dbErr, ok := err.(*pq.Error); ok {
					if pq.ErrorCode(tc.errorCode) != dbErr.Code {
						t.Fatalf("unexpected error code. want %v, got %v", tc.errorCode, dbErr.Code)
					}
					if tc.errorMessage != dbErr.Message {
						t.Fatalf("unexpected error message. want %v, got %v", tc.errorMessage, dbErr.Message)
					}
				}
				t.Skip("skipping test when expected error")
			}
			if user.Fullname != tc.arg.Fullname {
				t.Fatalf("unexpected fullname. want %v, got %v", tc.arg.Fullname, user.Fullname)
			}
			if user.Email != tc.arg.Email {
				t.Fatalf("unexpected email. want %v, got %v", tc.arg.Email, user.Email)
			}

			if user.ID == 0 {
				t.Fatalf("unexpected ID. got %d", user.ID)
			}
		})
	}
}
