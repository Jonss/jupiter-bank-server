package server

import (
	"fmt"
	"net/http"
)

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "alive and kicking")
	}
}

func (s *Server) Routes() {
	s.router.HandleFunc("/health", s.AppClientMiddleware(health())).Methods(http.MethodGet)
	s.router.HandleFunc("/sign-up", s.AppClientMiddleware(s.Signup())).Methods(http.MethodPost)
	s.router.HandleFunc("/sign-in", s.AppClientMiddleware(s.Authenticate())).Methods(http.MethodPost)
	s.router.HandleFunc("/test", s.PrivateRouteMiddleware(health())).Methods(http.MethodPost) // todo - reove it
}
