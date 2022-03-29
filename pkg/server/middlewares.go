package server

import (
	"context"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"net/http"
	"strings"
)

const authorization = "Authorization"
const userIDKey = "userID"

func (s *Server) AppClientMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get(authorization)
		if authorization == "" {
			rest.JsonResponse(w, http.StatusUnauthorized, rest.AuthorizationIsInvalid)
			return
		}
		splittedToken := strings.Split(authorization, " ")
		if len(splittedToken) < 2 {
			rest.JsonResponse(w, http.StatusUnauthorized, rest.AuthorizationIsInvalid)
			return
		}
		token := splittedToken[1]
		ok, err := s.basicAuthService.FetchAppClient(r.Context(), basic_auth.FetchAppClientParams{
			Token: token,
		})
		if err != nil {
			rest.JsonResponse(w, http.StatusInternalServerError, rest.UnexpectedError)
			return
		}
		if !ok {
			rest.JsonResponse(w, http.StatusUnauthorized, rest.Unauthorized)
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
			rest.JsonResponse(w, http.StatusUnauthorized, rest.Unauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), userIDKey, user.ID)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
