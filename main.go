package main

import (
	"log"
	"net/http"
)

func main() {
	server := LocationServer{store: &InMemoryLocationStore{
		locations: map[int]Location{
			1: {Name: "Mary", Id: 1},
		},
	}}
	log.Fatal(http.ListenAndServe(":5000", &server))
}
