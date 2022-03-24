package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"net/http"
	"strings"
)

func (s *Server) AppClientMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
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
		if err != nil || !ok {
			rest.JsonResponse(w, http.StatusUnauthorized, rest.Unauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func PrivateRouteMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}
