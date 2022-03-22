package main

import (
	"fmt"
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
	userDomain := user.NewUserDomain(q)
	srv := server.NewServer(router, userDomain)
	srv.Routes() // start routes

	fmt.Printf("Jupiter bank server running on [%s]. Env: %s", cfg.Port, cfg.Env)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, router))
}
