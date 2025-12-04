package main

import (
	"fmt"
	"net/http"
	"strings"
)

type InMemoryCustomerStore struct {
	customers []Customer
}

type InMemoryCustomerServer struct {
	store CustomerStore
}

func (i *InMemoryCustomerStore) GetCustomerBalance(name string) float64 {
	var balance = 123.00

	return balance
}

func (i *InMemoryCustomerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	customer := strings.TrimPrefix(r.URL.Path, "/customer/")
	balance := i.store.GetCustomerBalance(customer)

	fmt.Fprint(w, balance)
}

func (i *InMemoryCustomerStore) GetCustomers() Customers {
	return nil
}
