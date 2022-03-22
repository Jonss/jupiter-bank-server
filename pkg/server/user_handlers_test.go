package server

import (
	"bytes"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestSignup(t *testing.T) {
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
