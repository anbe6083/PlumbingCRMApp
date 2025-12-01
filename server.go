package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Customer struct {
	Name    string
	Balance string
}

func CustomerServer(w http.ResponseWriter, r *http.Request) {
	customer := strings.TrimPrefix(r.URL.Path, "/customer/")
	balance := GetCustomerBalance(customer)

	fmt.Fprint(w, balance)
}

func GetCustomerBalance(c string) string {
	if c == "Adam" {
		return "20000"
	}
	if c == "Mary" {
		return "10000"
	}

	return ""
}
