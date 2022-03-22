package server

import (
	"fmt"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive and kicking")
}

func (s *Server) Routes() {
	s.router.HandleFunc("/health", health).Methods(http.MethodGet)
	s.router.HandleFunc("/sign-up", s.userDomain.Signup()).Methods(http.MethodPost)
}
