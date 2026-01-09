package main

import (
	"log"
	"net/http"
)

func main() {
	server := LocationServer{store: &InMemoryLocationStore{}}
	log.Fatal(http.ListenAndServe(":5000", &server))
}
