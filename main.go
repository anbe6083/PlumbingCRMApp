package main

import (
	"log"
	"net/http"
)

func main() {
	store := &InMemoryCustomerStore{[]Customer{{Name: "mary", Balance: 10000}, {Name: "adam", Balance: 20000}}}
	server := &InMemoryCustomerServer{store}

	log.Fatal(http.ListenAndServe(":5000", server))
}
