package server

import (
	"bytes"
	"encoding/json"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup_Success(t *testing.T) {
	validator, _ := rest.NewValidator()

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
				"password": "123456789",
				"email": "jupiter.stein@jupiterbank.com"
			}
			`,
			userService:        &userServiceMock{},
			authorizationToken: "Basic banana",
			wantStatusCode:     http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), validator, tc.userService, &basicAuthMock{})
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
			r.Header.Add("Authorization", tc.authorizationToken)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, r)

			result := w.Result()
			gotStatusCode := result.StatusCode
			if tc.wantStatusCode != gotStatusCode {
				t.Fatalf("POST /sign-up . status code. want %d, got %d", tc.wantStatusCode, gotStatusCode)
			}
		})
	}
}

func TestSignup_Error(t *testing.T) {
	validator, _ := rest.NewValidator()

	testCases := []struct {
		name               string
		requestBody        string
		userService        user.Service
		wantStatusCode     int
		authorizationToken string
		wantErrorResponse  rest.ErrorResponses
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
			authorizationToken: "Basic banana",
			userService:       &userServiceMock{user.ErrUserExists},
			wantStatusCode:    http.StatusUnprocessableEntity,
			wantErrorResponse: rest.UserExists,
		},
		{
			name:           "should get error response when request body is empty",
			requestBody:    `{}`,
			authorizationToken: "Basic banana",
			userService:    &userServiceMock{},
			wantStatusCode: http.StatusBadRequest,
			wantErrorResponse: rest.NewErrorResponses(
				rest.NewValidationError("fullname is a required field"),
				rest.NewValidationError("email is a required field"),
				rest.NewValidationError("password is a required field"),
			),
		},
		{
			name: "should get error response when request body fields are invalid",
			requestBody: `{
				"fullname": "Eu",
				"email": "a-invalid-email#jupiterbank.com",
				"password": "12"
			}`,
			authorizationToken: "Basic banana",
			userService:    &userServiceMock{},
			wantStatusCode: http.StatusBadRequest,
			wantErrorResponse: rest.NewErrorResponses(
				rest.NewValidationError("fullname must be at least 3 characters in length"),
				rest.NewValidationError("email must be a valid email address"),
				rest.NewValidationError("password must be at least 6 characters in length"),
			),
		},
		{
			name: "should get error response when password is fields are invalid",
			requestBody: `{
				"fullname": "Jupiter Stein",
				"email": "a-valid-email@jupiterbank.com",
				"password": "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
			}`,
			authorizationToken: "Basic banana",
			userService:    &userServiceMock{},
			wantStatusCode: http.StatusBadRequest,
			wantErrorResponse: rest.NewErrorResponses(
				rest.NewValidationError("password must be at maximum 100 characters in length"),
			),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), validator, tc.userService, &basicAuthMock{})
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
			r.Header.Add("Authorization", tc.authorizationToken)
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, r)

			result := w.Result()
			gotStatusCode := result.StatusCode
			if tc.wantStatusCode != gotStatusCode {
				t.Fatalf("POST /sign-up . status code. want %d, got %d", tc.wantStatusCode, gotStatusCode)
			}

			resultBody, err := ioutil.ReadAll(result.Body)
			if err != nil {
				t.Fatal("unexpected error reading response body", err)
			}

			var errorResponse rest.ErrorResponses
			err = json.Unmarshal(resultBody, &errorResponse)
			if err != nil {
				t.Fatal("unexpected error unmarshalling response body", err)
			}

			for i, errResp := range errorResponse.Errors {
				if tc.wantErrorResponse.Errors[i].Message != errResp.Message {
					t.Fatalf("POST /sign-up. want error message %s, got %s", tc.wantErrorResponse.Errors[i].Message, errResp.Message)
				}
				if tc.wantErrorResponse.Errors[i].Code != errResp.Code {
					t.Fatalf("POST /sign-up. want error code %s, got %s", tc.wantErrorResponse.Errors[i].Code, errResp.Code)
				}
			}
		})
	}
}
