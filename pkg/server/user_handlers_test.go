package server

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServerSignup_Success(t *testing.T) {
	validator, _ := NewValidator()

	testCases := []struct {
		name               string
		requestBody        string
		authorizationToken string
		userService        user.Service
		wantStatusCode     int
	}{
		{
			name: "should signup successfully",
			requestBody: `
			{
				"fullname": "Jupiter Stein",
				"email": "jupiter.stein@jupiterbank.com",
				"password": "123456789"
			}
			`,
			userService:        &userServiceMock{},
			authorizationToken: "Basic banana",
			wantStatusCode:     http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), fakeConfig, validator, tc.userService, &basicAuthMock{}, &pasetoAuthMock{})
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/api/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
			r.Header.Add("Authorization", tc.authorizationToken)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, r)

			result := w.Result()
			gotStatusCode := result.StatusCode
			if tc.wantStatusCode != gotStatusCode {
				t.Fatalf("POST /api/sign-up . status code. want %d, got %d", tc.wantStatusCode, gotStatusCode)
			}
		})
	}
}

func TestServerSignup_Error(t *testing.T) {
	validator, _ := NewValidator()

	testCases := []struct {
		name               string
		requestBody        string
		userService        user.Service
		wantStatusCode     int
		authorizationToken string
		wantErrorResponse  ErrorResponses
	}{
		{
			name: "should not signup when user already exists",
			requestBody: `
			{
				"fullname": "Jupiter Stein",
				"password": "123456789",
				"email": "existing.user@jupiterbank.com"
			}
			`,
			authorizationToken: "Basic auth",
			userService:        &userServiceMock{err: user.ErrUserExists},
			wantStatusCode:     http.StatusUnprocessableEntity,
			wantErrorResponse:  UserExists,
		},
		{
			name:               "should get error response when request body is empty",
			requestBody:        `{}`,
			authorizationToken: "Basic banana",
			userService:        &userServiceMock{},
			wantStatusCode:     http.StatusBadRequest,
			wantErrorResponse: NewErrorResponses(
				NewValidationError("fullname is a required field"),
				NewValidationError("email is a required field"),
				NewValidationError("password is a required field"),
			),
		},
		{
			name: "should get error response when request body fields are invalid",
			requestBody: `{
				"fullname": "Eu",
				"email": "a-invalid-email#jupiterbank.com",
				"password": "12"
			}`,
			authorizationToken: "Basic auth",
			userService:        &userServiceMock{},
			wantStatusCode:     http.StatusBadRequest,
			wantErrorResponse: NewErrorResponses(
				NewValidationError("fullname must be at least 3 characters in length"),
				NewValidationError("email must be a valid email address"),
				NewValidationError("password must be at least 6 characters in length"),
			),
		},
		{
			name: "should get error response when password is fields are invalid",
			requestBody: `{
				"fullname": "Jupiter Stein",
				"email": "a-valid-email@jupiterbank.com",
				"password": "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
			}`,
			authorizationToken: "Basic auth",
			userService:        &userServiceMock{},
			wantStatusCode:     http.StatusBadRequest,
			wantErrorResponse: NewErrorResponses(
				NewValidationError("password must be at maximum 100 characters in length"),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), fakeConfig, validator, tc.userService, &basicAuthMock{}, &pasetoAuthMock{})
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/api/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
			r.Header.Add("Authorization", tc.authorizationToken)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, r)

			result := w.Result()
			gotStatusCode := result.StatusCode
			if tc.wantStatusCode != gotStatusCode {
				t.Fatalf("POST /api/sign-up . status code. want %d, got %d", tc.wantStatusCode, gotStatusCode)
			}

			resultBody, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatal("unexpected error reading response body", err)
			}

			var errorResponse ErrorResponses
			err = json.Unmarshal(resultBody, &errorResponse)
			if err != nil {
				t.Fatal("unexpected error unmarshalling response body", err)
			}

			for i, errResp := range errorResponse.Errors {
				if tc.wantErrorResponse.Errors[i].Message != errResp.Message {
					t.Fatalf("POST /api/sign-up. want error message %s, got %s", tc.wantErrorResponse.Errors[i].Message, errResp.Message)
				}
				if tc.wantErrorResponse.Errors[i].Code != errResp.Code {
					t.Fatalf("POST /api/sign-up. want error code %s, got %s", tc.wantErrorResponse.Errors[i].Code, errResp.Code)
				}
			}
		})
	}
}

func TestServer_Profile(t *testing.T) {
	validator, _ := NewValidator()

	testCases := []struct {
		name           string
		userService    user.Service
		wantStatusCode int
	}{
		{
			name:           "should get user",
			userService:    &userServiceMock{wantUser: &db.User{ID: 1}},
			wantStatusCode: http.StatusOK,
		},
		{
			name:           "should not get user when user does not exists",
			userService:    &userServiceMock{err: sql.ErrNoRows},
			wantStatusCode: http.StatusInternalServerError,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodGet, "/api/users", nil)
			w := httptest.NewRecorder()
			ctx := context.WithValue(r.Context(), userIDKey, "1")
			r.WithContext(ctx)
			srv := NewServer(mux.NewRouter(), fakeConfig, validator, tc.userService, &basicAuthMock{}, &pasetoAuthMock{})
			srv.Routes()

			srv.router.ServeHTTP(w, r)

			result := w.Result()
			if tc.wantStatusCode != result.StatusCode {
				t.Fatalf("POST /api/users. want %d, got %d", tc.wantStatusCode, result.StatusCode)
			}

		})
	}
}
