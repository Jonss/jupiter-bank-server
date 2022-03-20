package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jonss/jupiter-bank-server/pkg/config"
	"github.com/Jonss/jupiter-bank-server/pkg/db"
)

var q db.Queries
var cfg config.Config

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, cfg.Env)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive and kicking")
}

func configValue(w http.ResponseWriter, r *http.Request) {
	c, err := q.GetConfigByKey(r.Context(), db.GetConfigByKeyParams{Key: "a_key"})
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}
	fmt.Fprintf(w, "{key: %s, value: %s}", c.Key, c.Value)
}

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

	q = *db.New(conn)

	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/configs", configValue)
	fmt.Printf("Jupiter bank server running on [%s]. Env: %s", cfg.Port, cfg.Env)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}
