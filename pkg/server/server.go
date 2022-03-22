package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/gorilla/mux"
)

type Server struct {
	router      *mux.Router
	userService user.Service
}

func NewServer(r *mux.Router, ud user.Service) *Server {
	return &Server{router: r, userService: ud}
}
