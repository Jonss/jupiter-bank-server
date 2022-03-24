package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"github.com/gorilla/mux"
)

type Server struct {
	router      *mux.Router
	restValidator   *rest.Validator
	userService user.Service
}

func NewServer(r *mux.Router, v *rest.Validator, ud user.Service) *Server {
	return &Server{router: r, restValidator: v, userService: ud}
}
