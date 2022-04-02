package server

import (
	"fmt"
	"net/http"
)

func (s *Server) health() http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
		Env     string `json:"env"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		apiResponse(w, http.StatusOK, response{
			Message: "alive and kicking",
			Env:     s.config.Env,
		})
	}
}

func (s *Server) about() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Jupiter bank is under construction. v0.0.1")
	}
}

func (s *Server) Routes() {
	s.router.HandleFunc("/about", s.health()).Methods(http.MethodGet)
	s.router.HandleFunc("/health", s.AppClientMiddleware(s.health())).Methods(http.MethodGet)
	s.router.HandleFunc("/api/sign-up", s.AppClientMiddleware(s.Signup())).Methods(http.MethodPost)
	s.router.HandleFunc("/api/sign-in", s.AppClientMiddleware(s.Authenticate())).Methods(http.MethodPost)
	s.router.HandleFunc("/api/users", s.PrivateRouteMiddleware(s.Profile())).Methods(http.MethodGet)
}
