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
	testCases := []struct {
		name           string
		requestBody    string
		userService    user.Service
		wantStatusCode int
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
			userService:    &userServiceMock{},
			wantStatusCode: http.StatusCreated,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), tc.userService)
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
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
	testCases := []struct {
		name              string
		requestBody       string
		userService       user.Service
		wantStatusCode    int
		wantErrorResponse rest.ErrorResponses
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
			userService:       &userServiceMock{user.ErrUserExists},
			wantStatusCode:    http.StatusUnprocessableEntity,
			wantErrorResponse: rest.UserExists,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), tc.userService)
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer([]byte(tc.requestBody)))
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

			if errorResponse.Errors[0].Message != tc.wantErrorResponse.Errors[0].Message {
				t.Fatalf("POST /sign-up. want error message  %s, got %s", tc.wantErrorResponse.Errors[0].Message, errorResponse.Errors[0].Message)
			}

			if errorResponse.Errors[0].Code != tc.wantErrorResponse.Errors[0].Code {
				t.Fatalf("POST /sign-up. want error code %s, got %s", tc.wantErrorResponse.Errors[0].Code, errorResponse.Errors[0].Code)
			}
		})
	}
}
