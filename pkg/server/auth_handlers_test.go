package server

import (
	"bytes"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/paseto_auth"
	"github.com/gorilla/mux"
)

func TestAuthenticate(t *testing.T) {
	t.Parallel()
	validator, _ := NewValidator()

	testCases := []struct {
		name           string
		pasetoService  paseto_auth.Service
		requestBody    string
		wantStatusCode int
	}{
		{
			name:          "should authenticate - 200",
			pasetoService: &pasetoAuthMock{},
			requestBody: `
				{
					"email": "jupiter.stein@jupiterbank.com",
					"password": "123456789"
				}
			`,
			wantStatusCode: http.StatusOK,
		},
		{
			name:          "should not authenticate when user does not exist - 404",
			pasetoService: &pasetoAuthMock{err: sql.ErrNoRows},
			requestBody: `
				{
					"email": "jupiter.stein@jupiterbank.com",
					"password": "123456789"
				}
			`,
			wantStatusCode: http.StatusNotFound,
		},
		{
			name:           "should not authenticate when request body is empty - 400",
			pasetoService:  &pasetoAuthMock{},
			requestBody:    `{}`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:          "should not authenticate when email is invalid - 400",
			pasetoService: &pasetoAuthMock{},
			requestBody: `{
				"email": "invalid-email-format#email.com"
				"password": "123456789"
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
		{
			name:          "should not authenticate when password is invalid - 400",
			pasetoService: &pasetoAuthMock{},
			requestBody: `{
				"email": "valid@email.com"
				"password": ""
			}`,
			wantStatusCode: http.StatusBadRequest,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			srv := NewServer(mux.NewRouter(), fakeConfig, validator, &userServiceMock{}, &basicAuthMock{}, tc.pasetoService)
			srv.Routes()

			r := httptest.NewRequest(http.MethodPost, "/api/sign-in", bytes.NewBuffer([]byte(tc.requestBody)))
			r.Header.Add("Authorization", "Basic Auth")
			w := httptest.NewRecorder()
			srv.router.ServeHTTP(w, r)

			result := w.Result()
			if tc.wantStatusCode != result.StatusCode {
				t.Fatalf("/api/sign-in. want %d got %d", tc.wantStatusCode, result.StatusCode)
			}
		})
	}
}
