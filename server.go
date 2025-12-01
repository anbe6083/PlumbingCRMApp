package main

import (
	"fmt"
	"net/http"
	"strings"
)

const NotFoundStatus = http.StatusNotFound
const StatusAccepted = http.StatusAccepted

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

	switch r.Method {
	case http.MethodGet:
		c.showBalance(customer, w)
		return
	case http.MethodPost:
		c.processNewCustomer(w)
		return
	}

}

func (c *CustomerServer) showBalance(customer string, w http.ResponseWriter) {
	balance := c.store.GetCustomerBalance(customer)

	if balance == 0 {
		w.WriteHeader(NotFoundStatus)
	}
	fmt.Fprint(w, balance)
}

func (c *CustomerServer) processNewCustomer(w http.ResponseWriter) {
	w.WriteHeader(StatusAccepted)
}
