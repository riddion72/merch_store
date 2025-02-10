package main

import (
	"log"
	"net/http"
	"time"

	sw "./go"
)

const (
	shutdownTimeout = 5 * time.Second
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
