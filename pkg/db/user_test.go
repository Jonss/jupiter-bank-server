package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"testing"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func createUser(t *testing.T, uuid, fullname, email string) {
	insertQuery := fmt.Sprintf(`
	INSERT INTO users(
		external_id,
		fullname,
		email,
		password_hash
	) VALUES (
		'%s',
		'%s',
		'%s',
		'a_hashed_password'
	)`, uuid, fullname, email)

	_, err := testQueries.db.ExecContext(context.Background(), insertQuery)
	if err != nil {
		t.Fatalf("unexpected error inserting existing user. error=(%v)", err)
	}
}

func TestCreateUser(t *testing.T) {
	fullname := faker.FirstName() + " " + faker.LastName()
	createUser(t, uuid.NewString(), fullname, "existing.user@jupiterbank.com")

	uuid1, _ := uuid.NewUUID()
	uuid2, _ := uuid.NewUUID()
	testCases := []struct {
		name         string
		arg          CreateUserParams
		wantError    bool
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
				Fullname:   fullname,
				Email:      "existing.user@jupiterbank.com",
				ExternalID: uuid2,
			},
			wantError:    true,
			errorCode:    "23505",
			errorMessage: `duplicate key value violates unique constraint "users_email_key"`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			user, err := testQueries.CreateUser(context.Background(), tc.arg)
			if tc.wantError {
				if dbErr, ok := err.(*pq.Error); ok {
					if pq.ErrorCode(tc.errorCode) != dbErr.Code {
						t.Fatalf("unexpected error code. want %v, got %v", tc.errorCode, dbErr.Code)
					}
					if tc.errorMessage != dbErr.Message {
						t.Fatalf("unexpected error message. want %v, got %v", tc.errorMessage, dbErr.Message)
					}
				}
			} else {
				if user.Fullname != tc.arg.Fullname {
					t.Fatalf("unexpected fullname. want %v, got %v", tc.arg.Fullname, user.Fullname)
				}
				if user.Email != tc.arg.Email {
					t.Fatalf("unexpected email. want %v, got %v", tc.arg.Email, user.Email)
				}

				if user.ID == 0 {
					t.Fatalf("unexpected ID. got %d", user.ID)
				}
			}
		})
	}
}

func TestCheckUserExistsByEmail(t *testing.T) {
	existingEmail := faker.Email()
	fullname := faker.FirstName() + " " + faker.LastName()
	createUser(t, uuid.NewString(), fullname, existingEmail)

	testCases := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "should return false when user does not exists",
			arg:  "i-don-exist@jupiterbank.com",
			want: false,
		},
		{
			name: "should return true when user exists",
			arg:  existingEmail,
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := testQueries.CheckUserExistsByEmail(context.Background(), tc.arg)
			if err != nil {
				t.Fatal("unexpected error when check if user exists", err)
			}
			if tc.want != got {
				t.Fatalf("CheckUserExistsByEmail. want %t got %t", tc.want, got)
			}
		})
	}
}

func TestFetchUserByEmail(t *testing.T) {
	existingEmail := faker.Email()
	fullname := faker.FirstName() + " " + faker.LastName()
	newUuid := uuid.New()
	createUser(t, newUuid.String(), fullname, existingEmail)

	testCases := []struct {
		name    string
		email   string
		want    *User
		wantErr error
	}{
		{
			name:  "should get user",
			email: existingEmail,
			want:  &User{Email: existingEmail, ExternalID: newUuid, Fullname: fullname},
		},
		{
			name:    "should return true when user exists",
			email:   "user-does-not-exist@jupiterbank.com",
			want:    nil,
			wantErr: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := testQueries.FetchUserByEmail(context.Background(), tc.email)
			if err != nil && err != tc.wantErr {
				t.Fatalf("FetchUserByEmail. want err %v got %v", tc.wantErr, err)
			}

			if tc.want != nil {
				if tc.want.Email != got.Email {
					t.Fatalf("FetchUserByEmail. want email %v got %v", tc.want.Email, got.Email)
				}
				if tc.want.Fullname != got.Fullname {
					t.Fatalf("FetchUserByEmail. want fullname %v got %v", tc.want.Fullname, got.Fullname)
				}

				if tc.want.ExternalID != got.ExternalID {
					t.Fatalf("FetchUserByEmail. want externalID %v got %v", tc.want.ExternalID, got.ExternalID)
				}
			}
		})
	}
}
