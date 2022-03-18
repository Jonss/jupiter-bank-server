package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jonss/jupiter-bank-server/pkg/db"
)

// TODO: config using viper
const localPort = "8080"
const localMessage = "[local] Hello, world!"
const localDbSource = "postgres://user:pass@localhost:5445/jupiterbank?sslmode=disable"
const localDbName = "jupiterbank"

var message string
var port string
var dbSource string
var dbName string
var q db.Queries

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, message)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive and kicking ")
}

func configValue(w http.ResponseWriter, r *http.Request) {
	c, err := q.GetConfigByKey(r.Context(), db.GetConfigByKeyParams{Key: "a_key"})
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("{key: %s, value: %s}", c.Key, c.Value))
}

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = localPort
	}
	message = os.Getenv("APP_MESSAGE")
	if message == "" {
		message = localMessage
	}
	dbSource = os.Getenv("DB_SOURCE")
	if dbSource == "" {
		dbSource = localDbSource
	}

	dbName = os.Getenv("DB_Name")
	if dbName == "" {
		dbName = localDbName
	}

	conn, err := db.NewConnection(dbSource)
	if err != nil {
		panic(err)
	}
	err = db.Migrate(conn, dbName, "file://pkg/db/migrations")
	if err != nil {
		panic(err)
	}
	q = *db.New(conn)

	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/configs", configValue)

	fmt.Println("Jupiter bank server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
