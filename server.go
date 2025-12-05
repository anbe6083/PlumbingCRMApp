package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const NotFoundStatus = http.StatusNotFound
const StatusAccepted = http.StatusAccepted

type CustomerStore interface {
	GetCustomerBalance(name string) float64
	GetCustomers() Customers
	RecordNewCustomer(c Customer)
}

type CustomerServer struct {
	store CustomerStore
}

func (c *CustomerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/customers", http.HandlerFunc(c.customerBaseHandler))
	router.Handle("/customer/", http.HandlerFunc(c.customerHandler))
	router.ServeHTTP(w, r)
}

func (c *CustomerServer) showBalance(customer string, w http.ResponseWriter) {
	balance := c.store.GetCustomerBalance(customer)

	if balance == 0 {
		w.WriteHeader(NotFoundStatus)
	}
	fmt.Fprint(w, balance)
}

func (c *CustomerServer) processNewCustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	if r.Body == nil {
		log.Fatal("Request body is nil")
		return
	}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		fmt.Errorf("Problem decoding request body: %v", err)
	}
	c.store.RecordNewCustomer(customer)

	w.WriteHeader(StatusAccepted)
	json.NewEncoder(w).Encode(customer)
}

func (c *CustomerServer) customerHandler(w http.ResponseWriter, r *http.Request) {
	customer := strings.TrimPrefix(r.URL.Path, "/customer/")
	switch r.Method {
	case http.MethodGet:
		c.showBalance(customer, w)
		return
	case http.MethodPost:
		c.processNewCustomer(w, r)
		return
	}
}

func (c *CustomerServer) customerBaseHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(c.getAllCustomers())
	w.WriteHeader(http.StatusOK)
}

func (c *CustomerServer) getAllCustomers() []Customer {
	return c.store.GetCustomers()

}
