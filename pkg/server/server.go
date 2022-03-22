package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/gorilla/mux"
)

type Server struct {
	router     *mux.Router
	userDomain *user.UserDomain
}

func NewServer(r *mux.Router, ud *user.UserDomain) *Server {
	return &Server{router: r, userDomain: ud}
}
