package server

import (
	"github.com/Jonss/jupiter-bank-server/pkg/config"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/paseto_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"github.com/gorilla/mux"
)

type Server struct {
	router            *mux.Router
	config            config.Config
	restValidator     *rest.Validator
	userService       user.Service
	basicAuthService  basic_auth.Service
	pasetoAuthService paseto_auth.Service
}

func NewServer(
	r *mux.Router,
	cfg config.Config,
	v *rest.Validator,
	ud user.Service,
	bas basic_auth.Service,
	pas paseto_auth.Service,
) *Server {
	return &Server{
		router:            r,
		config:            cfg,
		restValidator:     v,
		userService:       ud,
		basicAuthService:  bas,
		pasetoAuthService: pas,
	}
}
