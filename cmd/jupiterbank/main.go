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

func main() {
	port = localPort
	if os.Getenv("APP_MESSAGE") == "" {
		message = localMessage
	}

	http.HandleFunc("/", hello)
	fmt.Println("Jupiter bank server running on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
