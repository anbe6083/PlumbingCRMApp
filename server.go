package main

import (
	"fmt"
	"net/http"
	"strings"
)

const NotFoundStatus = http.StatusNotFound

type Customer struct {
	Name    string
	Balance string
}

type CustomerStore interface {
	GetCustomerBalance(name string) int
}

type CustomerServer struct {
	store CustomerStore
}

func (c *CustomerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	customer := strings.TrimPrefix(r.URL.Path, "/customer/")
	balance := c.store.GetCustomerBalance(customer)

	if balance == 0 {
		w.WriteHeader(NotFoundStatus)
	}

	fmt.Fprint(w, balance)
}
