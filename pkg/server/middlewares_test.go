package server

import (
	"errors"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppClientMiddleware(t *testing.T) {
	validator, _ := rest.NewValidator()

	testCases := []struct {
		name             string
		authorization    string
		wantStatusCode   int
		basicAuthService basic_auth.Service
		wantErr          bool
	}{
		{
			name:             "should return OK when authorization code is valid",
			authorization:    "Basic auth",
			wantStatusCode:   http.StatusOK,
			basicAuthService: &basicAuthMock{},
		},
		{
			name:             "should return Unauthorized when authorization header is incomplete",
			authorization:    "Basic",
			wantStatusCode:   http.StatusUnauthorized,
			basicAuthService: &basicAuthMock{},
			wantErr:          true,
		},
		{
			name:             "should return Unauthorized when authorization header is empty",
			authorization:    "",
			wantStatusCode:   http.StatusUnauthorized,
			basicAuthService: &basicAuthMock{err: errors.New("unexpected error")},
			wantErr:          true,
		},
		{
			name:             "should return internal server error when baseAuth service returns an error",
			authorization:    "Basic valid-authorization",
			wantStatusCode:   http.StatusInternalServerError,
			basicAuthService: &basicAuthMock{err: errors.New("unexpected error")},
			wantErr:          true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), fakeConfig, validator, &userServiceMock{}, tc.basicAuthService, &pasetoAuthMock{})
			srv.Routes()

			r := httptest.NewRequest(http.MethodGet, "/health", nil)
			r.Header.Set("Authorization", tc.authorization)
			w := httptest.NewRecorder()

			srv.router.ServeHTTP(w, r)

			result := w.Result()

			if tc.wantStatusCode != result.StatusCode {
				t.Fatalf("AppClientMiddleware(). want %v got %v", tc.wantStatusCode, result.StatusCode)
			}
		})
	}
}
