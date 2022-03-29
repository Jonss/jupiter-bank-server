package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"net/http"
)

func (s *Server) health() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
		Env     string `json:"env"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		rest.JsonResponse(w, http.StatusOK, response{
			Message: "alive and kicking",
			Env:     s.config.Env,
		})
	}
}

func (s *Server) Routes() {
	s.router.HandleFunc("/health", s.AppClientMiddleware(s.health())).Methods(http.MethodGet)
	s.router.HandleFunc("/sign-up", s.AppClientMiddleware(s.Signup())).Methods(http.MethodPost)
	s.router.HandleFunc("/sign-in", s.AppClientMiddleware(s.Authenticate())).Methods(http.MethodPost)
}
