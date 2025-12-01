package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETCustomer(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/customer/Adam", nil)
	response := httptest.NewRecorder()

	CustomerServer(response, request)

	got := response.Body.String()
	want := "20000"

	if got != want {
		t.Errorf("Got %q, Wanted %q", got, want)
	}
}
