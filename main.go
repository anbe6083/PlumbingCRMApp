package main

import (
	"log"
	"net/http"
)

func main() {
	store := InMemoryCustomerStore{balances: map[string]int{"Mary": 10000, "Adam": 20000}}
	server := &InMemoryCustomerServer{&store}

	log.Fatal(http.ListenAndServe(":5000", server))
}
