package main

import (
	"fmt"
	"net/http"
	"strings"
)

type InMemoryCustomerStore struct {
	balances map[string]int
}

type InMemoryCustomerServer struct {
	store CustomerStore
}

func (i *InMemoryCustomerStore) GetCustomerBalance(name string) int {
	var balance = 123

	return balance
}

func (i *InMemoryCustomerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	customer := strings.TrimPrefix(r.URL.Path, "/customer/")
	balance := i.store.GetCustomerBalance(customer)

	fmt.Fprint(w, balance)
}
