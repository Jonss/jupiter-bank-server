package server

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"net/http"
	"strings"
)

const authorization = "Authorization"

func (s *Server) AppClientMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get(authorization)
		if authorization == "" {
			apiResponse(w, http.StatusUnauthorized, AuthorizationIsInvalid)
			return
		}
		splittedToken := strings.Split(authorization, " ")
		if len(splittedToken) < 2 {
			apiResponse(w, http.StatusUnauthorized, AuthorizationIsInvalid)
			return
		}
		token := splittedToken[1]
		ok, err := s.basicAuthService.FetchAppClient(r.Context(), basic_auth.FetchAppClientParams{
			Token: token,
		})
		if err != nil {
			apiResponse(w, http.StatusInternalServerError, UnexpectedError)
			return
		}
		if !ok {
			apiResponse(w, http.StatusUnauthorized, Unauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func (s *Server) PrivateRouteMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hex := r.Header.Get("x-auth-hex")
		token := r.Header.Get(authorization)
		user, err := s.pasetoAuthService.VerifyUser(r.Context(), token, hex)
		if err != nil {
			apiResponse(w, http.StatusUnauthorized, Unauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, user.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
