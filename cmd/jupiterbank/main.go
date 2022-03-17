package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const localPort = "8080"
const localMessage = "[local] Hello, world!"

var message string
var port string

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, message)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "alive and kicking ")
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

	http.HandleFunc("/", hello)
	http.HandleFunc("/health", health)

	fmt.Println("Jupiter bank server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
