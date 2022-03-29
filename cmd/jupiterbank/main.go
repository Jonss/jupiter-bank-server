package main

import (
	"fmt"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/basic_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/auth/paseto_auth"
	"github.com/Jonss/jupiter-bank-server/pkg/server/rest"
	"log"
	"net/http"

	"github.com/Jonss/jupiter-bank-server/pkg/config"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
	"github.com/Jonss/jupiter-bank-server/pkg/domain/user"
	"github.com/Jonss/jupiter-bank-server/pkg/server"
	"github.com/gorilla/mux"
)

var cfg config.Config

func main() {
	var err error
	cfg, err = config.LoadConfig(".")
	if err != nil {
		log.Fatalf("error getting config: error=(%v)", err)
	}

	conn, err := db.NewConnection(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(conn, cfg.DBName, cfg.DBMigrationPath)
	if err != nil {
		log.Fatal(err)
	}

	q := db.New(conn)

	router := mux.NewRouter()
	validator, err := rest.NewValidator()
	if err != nil {
		log.Fatal(err)
	}
	userService := user.NewUserService(q)
	basicAuthService := basic_auth.NewBasicAuthService(q)
	pasetoAuthService := paseto_auth.NewPasetoAuthService(userService)

	srv := server.NewServer(
		router,
		cfg,
		validator,
		userService,
		basicAuthService,
		pasetoAuthService,
	)
	srv.Routes() // start routes

	fmt.Println(fmt.Sprintf("Jupiter bank server running on [%s]. Env: [%s]", cfg.Port, cfg.Env))
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
